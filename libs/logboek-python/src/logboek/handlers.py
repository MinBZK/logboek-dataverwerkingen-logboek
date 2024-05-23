from logging import getLogger

from google.protobuf import timestamp_pb2
from grpc import RpcError, insecure_channel

from logboek import ProcessingOperation, ProcessingOperationHandler
from logboek.proto.logboek.v1 import logboek_pb2, logboek_pb2_grpc


logger = getLogger(__name__)


class GrpcProcessingOperationHandler(ProcessingOperationHandler):
    def __init__(self, logboek_endpoint: str):
        self._logboek_url = logboek_endpoint
        self._client = logboek_pb2_grpc.LogboekServiceStub(insecure_channel(logboek_endpoint))

    def on_end(self, op: ProcessingOperation) -> None:
        attributes = []
        for key, value in op._attributes.items():
            attributes.append(logboek_pb2.Attribute(key=key, value=value))

        pb_op = logboek_pb2.ProcessingOperation(
            trace_id=op._context.trace_id.to_bytes(16),
            operation_id=op._context.operation_id.to_bytes(8),
            parent_operation_id=op._parent_context.operation_id.to_bytes(8) if op._parent_context else None,
            name=op._name,
            start_time=_to_timestamp(op._start_time),
            end_time=_to_timestamp(op._end_time),
            status_code=op._status_code.value,
            attributes=attributes,
        )
        request = logboek_pb2.ExportOperationsRequest(operations=(pb_op,))

        try:
            self._client.Export(request)
        except RpcError as error:
            logger.error("Failed to export procesing operation, error: %s, details: %s", error.code(), error.details())
        except Exception:
            raise


def _to_timestamp(nano: int) -> timestamp_pb2.Timestamp:
    ts = timestamp_pb2.Timestamp()
    ts.FromNanoseconds(nano)
    return ts
