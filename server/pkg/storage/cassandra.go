package storage

import (
	"bytes"
	"context"

	"github.com/gocql/gocql"
)

const cassandraKeyspace = "logboek"

type CassandraStore struct {
	config  *gocql.ClusterConfig
	session *gocql.Session
}

func NewCassandra(servers ...string) (*CassandraStore, error) {
	config := gocql.NewCluster(servers...)

	return &CassandraStore{config: config}, nil
}

func (s *CassandraStore) Initialize() error {
	initSession, err := s.config.CreateSession()
	if err != nil {
		return err
	}

	schema, _ := schemaFS.ReadFile("schema/cassandra.cql")
	stmts := bytes.Split(schema, []byte{';'})
	for _, stmt := range stmts[:len(stmts)-1] {
		if err = initSession.Query(string(stmt)).Exec(); err != nil {
			return err
		}
	}

	s.config.Keyspace = cassandraKeyspace
	session, err := s.config.CreateSession()
	if err != nil {
		return err
	}

	s.session = session

	return nil
}

func (s *CassandraStore) Close() error {
	s.session.Close()
	return nil
}

func (s *CassandraStore) Write(ctx context.Context, op ProcessingOperation) error {
	const insertOperation = `
		INSERT INTO processing_operations (
			trace_id, operation_id, parent_operation_id,
			name, start_time, end_time, status_code,
			foreign_trace_id, foreign_operation_id, attributes
		) VALUES (
			?, ?, ?,
			?, ?, ?, ?,
			?, ?, ?
		)`

	var parentOperationID *string
	if !op.ParentOperationID.Empty() {
		parentOperationID = stringPtr(op.ParentOperationID.String())
	}

	var foreignTraceID, foreignOperationID *string
	if !op.ForeignTraceID.Empty() && !op.ForeignOperationID.Empty() {
		foreignTraceID = stringPtr(op.ForeignTraceID.String())
		foreignOperationID = stringPtr(op.ForeignOperationID.String())
	}

	query := s.session.Query(
		insertOperation,
		op.TraceID.String(),
		op.OperationID.String(),
		parentOperationID,
		op.Name,
		op.StartTime,
		op.EndTime,
		op.StatusCode.String(),
		foreignTraceID,
		foreignOperationID,
		op.Attributes,
	).WithContext(ctx)

	return query.Exec()
}

func stringPtr(v string) *string {
	return &v
}
