    SELECT
        ServiceName,
        avgMap(StatusCounts) AS avg_by_status,
        maxMap(StatusCounts) AS peak_by_status
    FROM otel_http_status_counts
    GROUP BY ServiceName;
