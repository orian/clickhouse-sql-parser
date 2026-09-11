CREATE TABLE region_performance
(
    region String,
    datacenter String,
    avg_response_time AggregateFunction(avg, UInt32)
)
ENGINE = AggregatingMergeTree()
ORDER BY (datacenter, region);

CREATE MATERIALIZED VIEW region_performance_mv
TO region_performance
AS SELECT
    region,
    datacenter,
    avgMergeState(avg_response_time) AS avg_response_time
FROM server_performance
GROUP BY region, datacenter;

-- datacenter level table and materialized view

CREATE TABLE datacenter_performance
(
    datacenter String,
    avg_response_time AggregateFunction(avg, UInt32)
)
ENGINE = AggregatingMergeTree()
ORDER BY datacenter;

CREATE MATERIALIZED VIEW datacenter_performance_mv
TO datacenter_performance
AS SELECT
      datacenter,
      avgMergeState(avg_response_time) AS avg_response_time
FROM region_performance
GROUP BY datacenter;
