from logboek import get_current_proccessing


class TraceContextPropegator:
    def inject(self, target: dict):
        current_proccessing = get_current_proccessing()
        context = current_proccessing._context
        if context is None:
            return

        value = f"00-{context.trace_id:032x}-{context.operation_id:016x}-01"

        target["traceparent"] = value
