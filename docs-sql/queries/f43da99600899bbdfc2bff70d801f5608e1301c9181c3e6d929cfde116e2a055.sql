ALTER TABLE otel_traces
ADD COLUMN PodName String
MATERIALIZED ResourceAttributes['k8s.pod.name']
