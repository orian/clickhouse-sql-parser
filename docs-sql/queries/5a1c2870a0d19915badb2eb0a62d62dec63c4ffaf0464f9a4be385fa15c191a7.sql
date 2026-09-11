CREATE MATERIALIZED VIEW events_snapshot_mv
REFRESH EVERY 10 SECOND APPEND TO events_snapshot
AS SELECT
    now() AS ts,
    uuid,
    sum(count) AS count
FROM events
GROUP BY ALL;
