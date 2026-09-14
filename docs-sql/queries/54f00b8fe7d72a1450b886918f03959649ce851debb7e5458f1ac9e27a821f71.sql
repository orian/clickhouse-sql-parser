CREATE TABLE otel_traces_backfill
(
    Timestamp DateTime64(9),
    ServiceName LowCardinality(String),
    StatusCode LowCardinality(String),
    Duration UInt64
)
ENGINE = Null;
