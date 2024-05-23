# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: logboek/proto/logboek/v1/logboek.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n&logboek/proto/logboek/v1/logboek.proto\x12\nlogboek.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"v\n\x17\x45xportOperationsRequest\x12&\n\x08resource\x18\x01 \x01(\x0b\x32\x14.logboek.v1.Resource\x12\x33\n\noperations\x18\x02 \x03(\x0b\x32\x1f.logboek.v1.ProcessingOperation\"\x1a\n\x18\x45xportOperationsResponse\")\n\x08Resource\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\"\xb1\x04\n\x13ProcessingOperation\x12\x10\n\x08trace_id\x18\x01 \x01(\x0c\x12\x14\n\x0coperation_id\x18\x02 \x01(\x0c\x12\x1b\n\x13parent_operation_id\x18\x03 \x01(\x0c\x12\x0c\n\x04name\x18\x04 \x01(\t\x12.\n\nstart_time\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08\x65nd_time\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12?\n\x0bstatus_code\x18\x07 \x01(\x0e\x32*.logboek.v1.ProcessingOperation.StatusCode\x12U\n\x11\x66oreign_operation\x18\x08 \x01(\x0b\x32:.logboek.v1.ProcessingOperation.ForeignProcessingOperation\x12)\n\nattributes\x18\t \x03(\x0b\x32\x15.logboek.v1.Attribute\x1aT\n\x1a\x46oreignProcessingOperation\x12\x10\n\x08trace_id\x18\x01 \x01(\x0c\x12\x14\n\x0coperation_id\x18\x02 \x01(\x0c\x12\x0e\n\x06\x65ntity\x18\x03 \x01(\t\"P\n\nStatusCode\x12\x17\n\x13STATUS_CODE_UNKNOWN\x10\x00\x12\x12\n\x0eSTATUS_CODE_OK\x10\x01\x12\x15\n\x11STATUS_CODE_ERROR\x10\x02\"\'\n\tAttribute\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t2g\n\x0eLogboekService\x12U\n\x06\x45xport\x12#.logboek.v1.ExportOperationsRequest\x1a$.logboek.v1.ExportOperationsResponse\"\x00\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'logboek.proto.logboek.v1.logboek_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_EXPORTOPERATIONSREQUEST']._serialized_start=87
  _globals['_EXPORTOPERATIONSREQUEST']._serialized_end=205
  _globals['_EXPORTOPERATIONSRESPONSE']._serialized_start=207
  _globals['_EXPORTOPERATIONSRESPONSE']._serialized_end=233
  _globals['_RESOURCE']._serialized_start=235
  _globals['_RESOURCE']._serialized_end=276
  _globals['_PROCESSINGOPERATION']._serialized_start=279
  _globals['_PROCESSINGOPERATION']._serialized_end=840
  _globals['_PROCESSINGOPERATION_FOREIGNPROCESSINGOPERATION']._serialized_start=674
  _globals['_PROCESSINGOPERATION_FOREIGNPROCESSINGOPERATION']._serialized_end=758
  _globals['_PROCESSINGOPERATION_STATUSCODE']._serialized_start=760
  _globals['_PROCESSINGOPERATION_STATUSCODE']._serialized_end=840
  _globals['_ATTRIBUTE']._serialized_start=842
  _globals['_ATTRIBUTE']._serialized_end=881
  _globals['_LOGBOEKSERVICE']._serialized_start=883
  _globals['_LOGBOEKSERVICE']._serialized_end=986
# @@protoc_insertion_point(module_scope)