SELECT Timestamp::Date AS day,
         count() AS c
FROM otel_logs
GROUP BY day
ORDER BY c DESC
