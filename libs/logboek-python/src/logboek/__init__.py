import random
from contextlib import contextmanager
from contextvars import ContextVar
from enum import Enum
from time import time_ns
from typing import Dict, Iterator, NamedTuple, Optional


class Resource(NamedTuple):
    name: str
    version: str


class ProcessingContext:
    def __init__(self, trace_id: int, operation_id: int):
        self._trace_id = trace_id
        self._operation_id = operation_id

    @property
    def trace_id(self) -> int:
        return self._trace_id

    @property
    def operation_id(self) -> int:
        return self._operation_id


class StatusCode(Enum):
    UNKNOWN = 0
    OK = 1
    ERROR = 2


class ProcessingOperationHandler:
    def on_end(self, op: "ProcessingOperation") -> None:
        pass


class ProcessingOperation:
    def __init__(
        self,
        name: str,
        resource: Resource,
        context: ProcessingContext,
        parent_context: Optional[ProcessingContext] = None,
        handler: ProcessingOperationHandler = ProcessingOperationHandler(),
    ):
        self._name = name
        self._resource = resource
        self._context = context
        self._parent_context = parent_context
        self._handler = handler
        self._status_code = StatusCode.UNKNOWN
        self._attributes: Dict[str, str] = {}

    def start(self) -> None:
        self._start_time = time_ns()

    def end(self) -> None:
        self._end_time = time_ns()
        self._handler.on_end(self)

    def set_status(self, status_code: StatusCode) -> None:
        self._status_code = status_code

    def set_attributes(self, attributes: Dict[str, str]) -> None:
        for key, value in attributes.items():
            self._attributes[key] = value

    def set_attribute(self, key: str, value: str) -> None:
        self.set_attributes({key: value})


_NONE_PROCESSING = ProcessingOperation("", None, None)
_CURRENT_PROCESSING = ContextVar("current-processing", default=_NONE_PROCESSING)


def get_current_proccessing() -> ProcessingOperation:
    return _CURRENT_PROCESSING.get()


class ProcessingOperator:
    def __init__(self, resource: Resource, handler: ProcessingOperationHandler) -> None:
        self._resource = resource
        self._handler = handler

    def start_proccessing(self, processing_name: str) -> ProcessingOperation:
        current_processing = get_current_proccessing()
        if current_processing is not _NONE_PROCESSING:
            parent_context = current_processing._context
            trace_id = parent_context.trace_id
        else:
            parent_context = None
            trace_id = self._generate_trace_id()

        context = ProcessingContext(trace_id=trace_id, operation_id=self._generate_operation_id())

        op = ProcessingOperation(processing_name, self._resource, context, parent_context, self._handler)
        op.start()

        return op

    @contextmanager
    def start_proccessing_as_current(self, processing_name: str) -> Iterator[ProcessingOperation]:
        op = self.start_proccessing(processing_name)
        try:
            token = _CURRENT_PROCESSING.set(op)
            try:
                yield op
            finally:
                _CURRENT_PROCESSING.reset(token)
        except Exception:
            op.set_status(StatusCode.ERROR)
            raise
        finally:
            op.end()

    def _generate_trace_id(self) -> int:
        return random.getrandbits(128)

    def _generate_operation_id(self) -> int:
        return random.getrandbits(64)


_PROCESSING_OPERATOR: Optional["ProcessingOperator"] = None


def init_processing_operator(
    resource: Resource, handler: Optional[ProcessingOperationHandler] = ProcessingOperationHandler()
):
    global _PROCESSING_OPERATOR
    if _PROCESSING_OPERATOR is not None:
        raise RuntimeError("Processing operator already initialized")

    _PROCESSING_OPERATOR = ProcessingOperator(resource, handler)


def get_processing_operator() -> ProcessingOperator:
    if _PROCESSING_OPERATOR is None:
        raise RuntimeError("Processing operator is not initialized")

    return _PROCESSING_OPERATOR
