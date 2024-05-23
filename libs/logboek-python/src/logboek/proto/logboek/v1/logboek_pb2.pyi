from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class ExportOperationsRequest(_message.Message):
    __slots__ = ("resource", "operations")
    RESOURCE_FIELD_NUMBER: _ClassVar[int]
    OPERATIONS_FIELD_NUMBER: _ClassVar[int]
    resource: Resource
    operations: _containers.RepeatedCompositeFieldContainer[ProcessingOperation]
    def __init__(self, resource: _Optional[_Union[Resource, _Mapping]] = ..., operations: _Optional[_Iterable[_Union[ProcessingOperation, _Mapping]]] = ...) -> None: ...

class ExportOperationsResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class Resource(_message.Message):
    __slots__ = ("name", "version")
    NAME_FIELD_NUMBER: _ClassVar[int]
    VERSION_FIELD_NUMBER: _ClassVar[int]
    name: str
    version: str
    def __init__(self, name: _Optional[str] = ..., version: _Optional[str] = ...) -> None: ...

class ProcessingOperation(_message.Message):
    __slots__ = ("trace_id", "operation_id", "parent_operation_id", "name", "start_time", "end_time", "status_code", "foreign_operation", "attributes")
    class StatusCode(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        STATUS_CODE_UNKNOWN: _ClassVar[ProcessingOperation.StatusCode]
        STATUS_CODE_OK: _ClassVar[ProcessingOperation.StatusCode]
        STATUS_CODE_ERROR: _ClassVar[ProcessingOperation.StatusCode]
    STATUS_CODE_UNKNOWN: ProcessingOperation.StatusCode
    STATUS_CODE_OK: ProcessingOperation.StatusCode
    STATUS_CODE_ERROR: ProcessingOperation.StatusCode
    class ForeignProcessingOperation(_message.Message):
        __slots__ = ("trace_id", "operation_id", "entity")
        TRACE_ID_FIELD_NUMBER: _ClassVar[int]
        OPERATION_ID_FIELD_NUMBER: _ClassVar[int]
        ENTITY_FIELD_NUMBER: _ClassVar[int]
        trace_id: bytes
        operation_id: bytes
        entity: str
        def __init__(self, trace_id: _Optional[bytes] = ..., operation_id: _Optional[bytes] = ..., entity: _Optional[str] = ...) -> None: ...
    TRACE_ID_FIELD_NUMBER: _ClassVar[int]
    OPERATION_ID_FIELD_NUMBER: _ClassVar[int]
    PARENT_OPERATION_ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    START_TIME_FIELD_NUMBER: _ClassVar[int]
    END_TIME_FIELD_NUMBER: _ClassVar[int]
    STATUS_CODE_FIELD_NUMBER: _ClassVar[int]
    FOREIGN_OPERATION_FIELD_NUMBER: _ClassVar[int]
    ATTRIBUTES_FIELD_NUMBER: _ClassVar[int]
    trace_id: bytes
    operation_id: bytes
    parent_operation_id: bytes
    name: str
    start_time: _timestamp_pb2.Timestamp
    end_time: _timestamp_pb2.Timestamp
    status_code: ProcessingOperation.StatusCode
    foreign_operation: ProcessingOperation.ForeignProcessingOperation
    attributes: _containers.RepeatedCompositeFieldContainer[Attribute]
    def __init__(self, trace_id: _Optional[bytes] = ..., operation_id: _Optional[bytes] = ..., parent_operation_id: _Optional[bytes] = ..., name: _Optional[str] = ..., start_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., end_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., status_code: _Optional[_Union[ProcessingOperation.StatusCode, str]] = ..., foreign_operation: _Optional[_Union[ProcessingOperation.ForeignProcessingOperation, _Mapping]] = ..., attributes: _Optional[_Iterable[_Union[Attribute, _Mapping]]] = ...) -> None: ...

class Attribute(_message.Message):
    __slots__ = ("key", "value")
    KEY_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    key: str
    value: str
    def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
