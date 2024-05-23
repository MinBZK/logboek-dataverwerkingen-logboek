# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from logboek.proto.logboek.v1 import logboek_pb2 as logboek_dot_proto_dot_logboek_dot_v1_dot_logboek__pb2

GRPC_GENERATED_VERSION = '1.63.0'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.65.0'
SCHEDULED_RELEASE_DATE = 'June 25, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in logboek/proto/logboek/v1/logboek_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class LogboekServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Export = channel.unary_unary(
                '/logboek.v1.LogboekService/Export',
                request_serializer=logboek_dot_proto_dot_logboek_dot_v1_dot_logboek__pb2.ExportOperationsRequest.SerializeToString,
                response_deserializer=logboek_dot_proto_dot_logboek_dot_v1_dot_logboek__pb2.ExportOperationsResponse.FromString,
                _registered_method=True)


class LogboekServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Export(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_LogboekServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Export': grpc.unary_unary_rpc_method_handler(
                    servicer.Export,
                    request_deserializer=logboek_dot_proto_dot_logboek_dot_v1_dot_logboek__pb2.ExportOperationsRequest.FromString,
                    response_serializer=logboek_dot_proto_dot_logboek_dot_v1_dot_logboek__pb2.ExportOperationsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'logboek.v1.LogboekService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class LogboekService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Export(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/logboek.v1.LogboekService/Export',
            logboek_dot_proto_dot_logboek_dot_v1_dot_logboek__pb2.ExportOperationsRequest.SerializeToString,
            logboek_dot_proto_dot_logboek_dot_v1_dot_logboek__pb2.ExportOperationsResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
