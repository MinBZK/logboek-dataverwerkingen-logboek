package logboek

import "context"

type ctxKey int

var currentProcessingKey ctxKey

type ProcessingContext struct {
	traceID     TraceID
	operationID OperationID
	foreign     bool
}

func (p ProcessingContext) TraceID() TraceID {
	return p.traceID
}

func (p ProcessingContext) OperationID() OperationID {
	return p.operationID
}

func (p ProcessingContext) IsValid() bool {
	return !p.traceID.Empty() && !p.operationID.Empty()
}

func ContextWithProcessingOperation(parent context.Context, op *ProcessingOperation) context.Context {
	return context.WithValue(parent, currentProcessingKey, op)
}

func ContexFromRemoteProcessingOperation() context.Context {
	return context.Background()
}

var emptyProcessingOperation = ProcessingOperation{}

func ProcessingOperationFromContext(ctx context.Context) *ProcessingOperation {
	if ctx == nil {
		return &emptyProcessingOperation
	}

	op, ok := ctx.Value(currentProcessingKey).(*ProcessingOperation)
	if !ok {
		return &emptyProcessingOperation
	}

	return op
}
