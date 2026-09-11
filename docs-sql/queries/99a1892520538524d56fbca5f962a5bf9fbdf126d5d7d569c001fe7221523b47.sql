SELECT Timestamp::Date AS day,
        count() AS c
FROM otel_logs_archive
GROUP BY day
ORDER BY c DESC
