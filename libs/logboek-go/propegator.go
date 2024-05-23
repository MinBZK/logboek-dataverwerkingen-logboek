package logboek

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
)

const (
	traceparentHeader = "traceparent"
)

type TraceContextPropegator struct{}

func (p TraceContextPropegator) Extract(ctx context.Context, h http.Header) context.Context {
	tp := h.Get(traceparentHeader)
	if tp == "" {
		return ctx
	}

	parts := strings.SplitN(tp, "-", 4)
	version := parts[0]
	if len(version) != 2 {
		return ctx
	}

	traceID, err := hex.DecodeString(parts[1])
	if err != nil || len(traceID) != 16 {
		return ctx
	}

	parentID, err := hex.DecodeString(parts[2])
	if err != nil || len(parentID) != 8 {
		return ctx
	}

	op := ProcessingOperation{
		context: ProcessingContext{
			traceID:     TraceID(traceID),
			operationID: OperationID(parentID),
			foreign:     true,
		},
	}

	return ContextWithProcessingOperation(ctx, &op)
}

func (p TraceContextPropegator) Inject(ctx context.Context, h http.Header) {
	op := ProcessingOperationFromContext(ctx)
	if !op.context.IsValid() {
		return
	}

	v := fmt.Sprintf("00-%x-%x-01", op.context.traceID[:], op.context.operationID[:])

	h.Set(traceparentHeader, v)
}
