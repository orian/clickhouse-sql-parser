CREATE TABLE otel_traces_1m
(
    `Timestamp` DateTime,
    `ServiceName` LowCardinality(String),
    `StatusCode` LowCardinality(String),
    `count` SimpleAggregateFunction(sum, UInt64),
    `avg__Duration` AggregateFunction(avg, UInt64),
    `max__Duration` SimpleAggregateFunction(max, Int64),
    `quantiles__Duration` AggregateFunction(quantiles(0.95, 0.99), Int64)
)
ENGINE = AggregatingMergeTree
ORDER BY (Timestamp, ServiceName, StatusCode);
