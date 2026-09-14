        SELECT
            ServiceName,
            sumMap(StatusCounts) AS total_by_status
        FROM otel_http_status_counts
        GROUP BY ServiceName;
