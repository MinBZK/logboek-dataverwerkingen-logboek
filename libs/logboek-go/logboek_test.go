package logboek_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go"
)

func TestStartProcessingParentContext(t *testing.T) {
	ctx := context.Background()

	operator := logboek.NewProcessingOperator(nil)
	ctx, a := operator.StartProcessing(ctx, "a")
	_, b := operator.StartProcessing(ctx, "b")

	pc := a.GetProcessingContext()
	ppc := b.GetParentProcessingContext()

	assert.Equal(t, pc.TraceID(), ppc.TraceID())
	assert.Equal(t, pc.OperationID(), ppc.OperationID())
}
