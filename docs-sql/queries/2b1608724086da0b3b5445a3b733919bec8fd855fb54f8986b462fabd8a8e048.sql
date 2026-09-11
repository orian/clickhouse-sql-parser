SELECT *
FROM prometheusQuery(
    prometheus.metrics,
    'rate(http_requests_total[5m])',
    now()
);
