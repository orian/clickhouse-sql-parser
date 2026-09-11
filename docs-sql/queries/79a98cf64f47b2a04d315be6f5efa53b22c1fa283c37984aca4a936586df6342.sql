    CREATE MATERIALIZED VIEW mv_events_rollup_1d
    TO events_rollup_1d
    AS
    SELECT
        toDate(event_time) AS bucket_start,
        country,
        event_type,
        uniqExactState(user_id),
        sumState(value),
        avgState(value),
        countState()
    FROM events_raw
    GROUP BY ALL;
