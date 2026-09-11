CREATE MATERIALIZED VIEW mv_events_rollup_1h
TO events_rollup_1h
AS
SELECT
    toStartOfHour(event_time) AS bucket_start,
    country,
    event_type,
    uniqExactState(user_id)   AS users_uniq,
    sumState(value)           AS value_sum,
    avgState(value)           AS value_avg,
    countState()              AS events_count
FROM events_raw
GROUP BY bucket_start, country, event_type;
