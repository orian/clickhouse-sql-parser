ALTER TABLE otel_traces
ADD INDEX idx_pod_name PodName
TYPE bloom_filter(0.01)
GRANULARITY 1
