SELECT
    region,
    datacenter,
    avgMerge(avg_response_time) AS avg_response_ms
FROM region_performance
GROUP BY region, datacenter
ORDER BY datacenter, region;
