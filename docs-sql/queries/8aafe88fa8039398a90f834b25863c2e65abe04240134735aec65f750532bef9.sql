SELECT * FROM prometheusQuery(mytable, 'rate(http_requests{job="prometheus"}[10m])[1h:10m]', now())
