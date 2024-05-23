package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/mattn/go-sqlite3"
)

type SqliteStore struct {
	db *sql.DB

	insertOperationStmt *sql.Stmt
	insertAttributeStmt *sql.Stmt
}

func NewSqlite(dbPath string) (*SqliteStore, error) {
	version, _, _ := sqlite3.Version()
	log.Printf("Using SQLite version: %s", version)

	dir, _ := path.Split(dbPath)
	if dir != "" {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.Mkdir(dir, 0750); err != nil {
				return nil, err
			}
		}
	}

	dsn := fmt.Sprintf("file:%s?cache=shared", dbPath)
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &SqliteStore{db: db}, nil
}

func (s *SqliteStore) Initialize() error {
	schema, _ := schemaFS.ReadFile("schema/sqlite.sql")

	if _, err := s.db.Exec(string(schema)); err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	var err error

	const insertOperation = `
		INSERT INTO processing_operations (
			trace_id, operation_id, parent_operation_id,
			name, start_time, end_time, status_code,
			foreign_trace_id, foreign_operation_id
		) VALUES (
			?, ?, ?,
			?, ?, ?, ?,
			?, ?
		)`
	s.insertOperationStmt, err = s.db.Prepare(insertOperation)
	if err != nil {
		return fmt.Errorf("failed to prepeare insert statement: %w", err)
	}

	s.insertAttributeStmt, err = s.db.Prepare("INSERT INTO processing_operation_attributes (trace_id, operation_id, key, value) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepeare insert statement: %w", err)
	}

	return nil
}

func (s *SqliteStore) Close() error {
	return s.db.Close()
}

func (s *SqliteStore) Write(ctx context.Context, op ProcessingOperation) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	traceID := op.TraceID.String()
	operationID := op.OperationID.String()

	var parentOperationID sql.NullString
	if !op.ParentOperationID.Empty() {
		parentOperationID.String = op.ParentOperationID.String()
		parentOperationID.Valid = true
	}

	var foreignTraceID, foreignOperationID sql.NullString
	if !op.ForeignTraceID.Empty() && !op.ForeignOperationID.Empty() {
		foreignTraceID.String, foreignTraceID.Valid = op.ForeignTraceID.String(), true
		foreignOperationID.String, foreignOperationID.Valid = op.ForeignOperationID.String(), true
	}

	_, err = tx.StmtContext(ctx, s.insertOperationStmt).ExecContext(
		ctx,
		traceID,
		operationID,
		parentOperationID,
		op.Name,
		op.StartTime.Format(time.RFC3339Nano),
		op.EndTime.Format(time.RFC3339Nano),
		op.StatusCode.String(),
		foreignTraceID,
		foreignOperationID,
	)
	if err != nil {
		return err
	}

	insertAttributeStmt := tx.StmtContext(ctx, s.insertAttributeStmt)

	for _, attribute := range op.Attributes {
		_, err = insertAttributeStmt.ExecContext(
			ctx,
			traceID,
			operationID,
			attribute.Key,
			attribute.Value,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()

	return err
}
