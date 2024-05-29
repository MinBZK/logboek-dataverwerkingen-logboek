CREATE KEYSPACE IF NOT EXISTS logboek WITH replication = {
    'class': 'SimpleStrategy', 'replication_factor': 1
};

CREATE TYPE IF NOT EXISTS logboek.attribute (
    key text,
    value text
);

CREATE TABLE IF NOT EXISTS logboek.processing_operations (
    trace_id text,
    operation_id text,
    parent_operation_id text,
    name text,
    start_time timestamp,
    end_time timestamp,
    status_code text,
    foreign_trace_id text,
    foreign_operation_id text,
    attributes list<frozen<attribute>>,
    resource_name text,
    resource_version text,

    PRIMARY KEY (trace_id, operation_id)
);
