SELECT
    page_id,
    page_name,
    avgMerge(avg_response_time) AS average_response_time_ms
FROM page_performance
GROUP BY page_id, page_name
ORDER BY page_id;
