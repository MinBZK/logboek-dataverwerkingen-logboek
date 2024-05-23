package logboek

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go/attribute"
)

type TraceID [16]byte

var emptyTraceID TraceID

func (id TraceID) Empty() bool {
	return bytes.Equal(id[:], emptyTraceID[:])
}

func (id TraceID) String() string {
	return fmt.Sprintf("%032x", id[:])
}

func TraceIDFromBytes(data []byte) (TraceID, error) {
	if l := len(data); l > 16 {
		return TraceID{}, fmt.Errorf("invalid length: %d", l)
	}

	return TraceID(data), nil
}

type OperationID [8]byte

var emptyOperationID OperationID

func (id OperationID) Empty() bool {
	return bytes.Equal(id[:], emptyOperationID[:])
}

func (id OperationID) String() string {
	return fmt.Sprintf("%016x", id[:])
}

func OperationIDFromBytes(data []byte) (OperationID, error) {
	if l := len(data); l > 8 {
		return OperationID{}, fmt.Errorf("invalid length: %d", l)
	}

	return OperationID(data), nil
}

type StatusCode int

const (
	StatusCodeUnknown StatusCode = iota
	StatusCodeOK
	StatusCodeError
)

type ProcessingOperation struct {
	name          string
	context       ProcessingContext
	parentContext ProcessingContext

	startTime  time.Time
	endTime    time.Time
	statusCode StatusCode
	attributes []attribute.Attribute

	handler ProcessingOperationHandler
}

func (op *ProcessingOperation) Start() {
	op.startTime = time.Now()
}

func (op *ProcessingOperation) End() {
	op.endTime = time.Now()
	op.handler.OnEnd(op)
}

func (op *ProcessingOperation) GetProcessingContext() ProcessingContext {
	return op.context
}

func (op *ProcessingOperation) GetParentProcessingContext() ProcessingContext {
	return op.parentContext
}

func (op *ProcessingOperation) SetStatus(c StatusCode) {
	op.statusCode = c
}

func (op *ProcessingOperation) SetAttributes(attrs ...attribute.Attribute) {
	op.attributes = append(op.attributes, attrs...)
}

type ProcessingOperator struct {
	handler ProcessingOperationHandler
}

func (o *ProcessingOperator) StartProcessing(ctx context.Context, processingName string) (context.Context, *ProcessingOperation) {
	var (
		parentContext ProcessingContext
		traceID       TraceID
		operationID   OperationID
	)

	if parent := ProcessingOperationFromContext(ctx); parent.context.IsValid() {
		parentContext = parent.context
	}

	if parentContext.IsValid() && !parentContext.foreign {
		traceID = parentContext.traceID
	} else {
		_, _ = rand.Read(traceID[:])
	}

	_, _ = rand.Read(operationID[:])

	op := &ProcessingOperation{
		name: processingName,
		context: ProcessingContext{
			traceID:     traceID,
			operationID: operationID,
		},
		parentContext: parentContext,
		handler:       o.handler,
	}

	ctx = ContextWithProcessingOperation(ctx, op)

	op.Start()

	return ctx, op
}

func NewProcessingOperator(handler ProcessingOperationHandler) *ProcessingOperator {
	return &ProcessingOperator{
		handler: handler,
	}
}
