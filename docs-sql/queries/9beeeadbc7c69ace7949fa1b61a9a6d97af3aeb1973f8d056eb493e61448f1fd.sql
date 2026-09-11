SELECT
  metric AS name,
  toFloat64(value) AS value,
  description AS help,
  labels,
  'gauge' AS type
FROM system.dimensional_metrics
FORMAT Prometheus
