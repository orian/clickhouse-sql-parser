CREATE TABLE raw_server_metrics
(
    timestamp DateTime DEFAULT now(),
    server_id UInt32,
    region String,
    datacenter String,
    response_time_ms UInt32
)
ENGINE = MergeTree()
ORDER BY (region, server_id, timestamp);
