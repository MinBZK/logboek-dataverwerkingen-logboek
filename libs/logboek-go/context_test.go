package logboek_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go"
)

func TestProcessingOperationFromContext(t *testing.T) {
	ctx := context.Background()

	operator := logboek.NewProcessingOperator(logboek.Resource{}, nil)
	ctx, op := operator.StartProcessing(ctx, "test")

	assert.Same(t, op, logboek.ProcessingOperationFromContext(ctx))
}
