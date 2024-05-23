package logboek

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	proto_v1 "github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go/proto/logboek/v1"
)

type ProcessingOperationHandler interface {
	OnEnd(op *ProcessingOperation)
}

type GRPCProcessingOperationHandler struct {
	client proto_v1.LogboekServiceClient
}

func (h *GRPCProcessingOperationHandler) OnEnd(op *ProcessingOperation) {
	out := proto_v1.ProcessingOperation{
		TraceId:     op.context.traceID[:],
		OperationId: op.context.operationID[:],
		Name:        op.name,
		StartTime:   timestamppb.New(op.startTime),
		EndTime:     timestamppb.New(op.endTime),
	}

	if op.parentContext.IsValid() {
		if op.parentContext.foreign {
			out.ForeignOperation = &proto_v1.ProcessingOperation_ForeignProcessingOperation{
				TraceId:     op.parentContext.traceID[:],
				OperationId: op.parentContext.operationID[:],
			}
		} else {
			out.ParentOperationId = op.parentContext.operationID[:]
		}
	}

	switch op.statusCode {
	case StatusCodeOK:
		out.StatusCode = proto_v1.ProcessingOperation_STATUS_CODE_OK
	case StatusCodeError:
		out.StatusCode = proto_v1.ProcessingOperation_STATUS_CODE_ERROR
	default:
		out.StatusCode = proto_v1.ProcessingOperation_STATUS_CODE_UNKNOWN
	}

	for _, attribute := range op.attributes {
		out.Attributes = append(out.Attributes, &proto_v1.Attribute{
			Key:   attribute.Key,
			Value: attribute.Value,
		})
	}

	req := proto_v1.ExportOperationsRequest{
		Operations: []*proto_v1.ProcessingOperation{&out},
	}

	ctx := context.Background()
	_, _ = h.client.Export(ctx, &req)
}

func NewGRPCProcessingOperationHandler(ctx context.Context, endpoint string) (*GRPCProcessingOperationHandler, error) {
	cc, err := grpc.DialContext(ctx, endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	h := GRPCProcessingOperationHandler{
		client: proto_v1.NewLogboekServiceClient(cc),
	}

	return &h, nil
}
