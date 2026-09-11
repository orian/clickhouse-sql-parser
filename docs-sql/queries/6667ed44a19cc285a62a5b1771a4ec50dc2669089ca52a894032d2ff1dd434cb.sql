ALTER TABLE otel_traces
ADD INDEX idx_kafka_offset KafkaOffset TYPE minmax GRANULARITY 1
