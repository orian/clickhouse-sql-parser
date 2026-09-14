    SELECT
        server_id,
        region,
        avgMerge(avg_response_time) AS avg_response_ms
    FROM server_performance
    GROUP BY server_id, region
    ORDER BY region, server_id;
