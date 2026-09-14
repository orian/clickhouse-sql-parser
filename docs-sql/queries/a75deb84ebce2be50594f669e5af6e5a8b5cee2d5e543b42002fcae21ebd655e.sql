ALTER TABLE otel_v2.otel_traces
MATERIALIZE INDEX idx_kafka_offset
IN PARTITION '2026-01-02';
