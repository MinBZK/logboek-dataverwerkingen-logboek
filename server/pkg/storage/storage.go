package storage

import (
	"context"
	"embed"
	"time"

	"github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go"
)

type Store interface {
	Initialize() error
	Close() error
	Write(ctx context.Context, op ProcessingOperation) error
}

//go:embed schema/*.*
var schemaFS embed.FS

type StatusCode int

const (
	StatuCodeUnknown StatusCode = iota
	StatusCodeOK
	StatusCodeError
)

func (s StatusCode) String() string {
	switch s {
	case StatuCodeUnknown:
		return "UNKNOWN"
	case StatusCodeOK:
		return "OK"
	case StatusCodeError:
		return "ERROR"
	}
	return ""
}

type Attribute struct {
	Key   string `cql:"key"`
	Value string `cql:"value"`
}

type ProcessingOperation struct {
	TraceID            logboek.TraceID
	OperationID        logboek.OperationID
	ParentOperationID  logboek.OperationID
	Name               string
	StartTime          time.Time
	EndTime            time.Time
	StatusCode         StatusCode
	ForeignTraceID     logboek.TraceID
	ForeignOperationID logboek.OperationID
	Attributes         []Attribute
	Resource           Resource
}

type Resource struct {
	Name    string
	Version string
}
