ALTER TABLE otel_traces
ADD COLUMN KafkaOffset UInt64
MATERIALIZED toUInt64(SpanAttributes['messaging.kafka.offset'])
