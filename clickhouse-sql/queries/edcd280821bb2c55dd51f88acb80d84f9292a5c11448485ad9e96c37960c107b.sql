
        CREATE MATERIALIZED VIEW mv_second TO mv_dst
        AS SELECT 0 AS key, value AS value FROM dst;
    