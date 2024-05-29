package server

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go"
	proto_v1 "github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go/proto/logboek/v1"

	"github.com/MinBZK/logboek-dataverwerkingen-logboek/server/pkg/storage"
)

type LogboekService struct {
	proto_v1.UnimplementedLogboekServiceServer

	store storage.Store
}

func (s LogboekService) Export(ctx context.Context, in *proto_v1.ExportOperationsRequest) (resp *proto_v1.ExportOperationsResponse, err error) {
	if in.Resource == nil {
		return nil, status.Errorf(codes.InvalidArgument, "missing resource")
	}

	resource := storage.Resource{
		Name:    in.Resource.Name,
		Version: in.Resource.Version,
	}

	for _, reqOp := range in.Operations {
		log.Printf("name=%s, root=%t, status_code=%s", reqOp.GetName(), reqOp.GetParentOperationId() == nil, reqOp.GetStatusCode())
		resp, err = s.handleOp(ctx, resource, reqOp)
		if err != nil {
			break
		}
	}

	return resp, err
}

func (s LogboekService) handleOp(ctx context.Context, resource storage.Resource, reqOp *proto_v1.ProcessingOperation) (*proto_v1.ExportOperationsResponse, error) {
	traceID, err := logboek.TraceIDFromBytes(reqOp.TraceId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "trace_id: %v", err)
	}

	operationID, err := logboek.OperationIDFromBytes(reqOp.OperationId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "operation_id: %v", err)
	}

	var parentOperationID logboek.OperationID
	if reqOp.ParentOperationId != nil {
		parentOperationID, err = logboek.OperationIDFromBytes(reqOp.ParentOperationId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "parent_operation_id: %v", err)
		}
	}

	var statusCode storage.StatusCode
	switch reqOp.StatusCode {
	case proto_v1.ProcessingOperation_STATUS_CODE_OK:
		statusCode = storage.StatusCodeOK
	case proto_v1.ProcessingOperation_STATUS_CODE_ERROR:
		statusCode = storage.StatusCodeError
	}

	var (
		foreignTraceID     logboek.TraceID
		foreignOperationID logboek.OperationID
	)

	if foreignOp := reqOp.ForeignOperation; foreignOp != nil {
		if traceID, operationID := foreignOp.GetTraceId(), foreignOp.GetOperationId(); traceID != nil && operationID != nil {
			foreignTraceID, err = logboek.TraceIDFromBytes(traceID)
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "foreign_operation.trace_id: %v", err)
			}
			foreignOperationID, err = logboek.OperationIDFromBytes(operationID)
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "foreign_operation.operation_id: %v", err)
			}
		}
	}

	attributes := make([]storage.Attribute, 0, len(reqOp.Attributes))
	for _, attribute := range reqOp.Attributes {
		attributes = append(attributes, storage.Attribute{
			Key:   attribute.Key,
			Value: attribute.Value,
		})
	}

	op := storage.ProcessingOperation{
		TraceID:            traceID,
		OperationID:        operationID,
		ParentOperationID:  parentOperationID,
		Name:               reqOp.Name,
		StartTime:          reqOp.StartTime.AsTime(),
		EndTime:            reqOp.EndTime.AsTime(),
		StatusCode:         statusCode,
		ForeignTraceID:     foreignTraceID,
		ForeignOperationID: foreignOperationID,
		Attributes:         attributes,
		Resource:           resource,
	}

	err = s.store.Write(ctx, op)
	if err != nil {
		log.Printf("failed to write to store: %v", err)
		return nil, status.Error(codes.Internal, "failed to store")
	}

	return &proto_v1.ExportOperationsResponse{}, nil
}
