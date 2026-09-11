    SELECT
        ServiceName,
        sumMapIf(StatusCounts, StatusCounts['5xx'] > 0) AS totals_in_error_windows
    FROM otel_http_status_counts
    GROUP BY ServiceName;
