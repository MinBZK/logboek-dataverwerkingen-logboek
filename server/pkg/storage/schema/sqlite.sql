CREATE TABLE IF NOT EXISTS processing_operations (
    trace_id CHAR(32) NOT NULL,
    operation_id CHAR(16) NOT NULL,
    parent_operation_id CHAR(16),
    name VARCHAR NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    status_code VARCHAR(5) NOT NULL,
    foreign_trace_id CHAR(32),
    foreign_operation_id CHAR(16),
    resource_name VARCHAR,
    resource_version VARCHAR,

    PRIMARY KEY (trace_id, operation_id)
) WITHOUT ROWID;

CREATE TABLE IF NOT EXISTS processing_operation_attributes (
    trace_id CHAR(32) NOT NULL,
    operation_id CHAR(16) NOT NULL,
    key VARCHAR NOT NULL,
    value VARCHAR NOT NULL,

    PRIMARY KEY (trace_id, operation_id, key),
    FOREIGN KEY (trace_id, operation_id) REFERENCES processing_operations (trace_id, operation_id)
) WITHOUT ROWID;
