CREATE MATERIALIZED VIEW page_performance_mv
TO page_performance
AS SELECT
    page_id,
    page_name,
    avgState(response_time_ms) AS avg_response_time  -- Usando o combinator -State
FROM raw_page_views
GROUP BY page_id, page_name;
