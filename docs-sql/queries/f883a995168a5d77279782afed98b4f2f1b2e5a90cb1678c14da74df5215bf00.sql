        SELECT
            bucket_start,
            country,
            event_type,
            uniqExactMerge(users_uniq) AS users,
            sumMerge(value_sum)        AS value_sum,
            avgMerge(value_avg)        AS value_avg,
            countMerge(events_count)   AS events
        FROM events_rollup_1h
        WHERE bucket_start >= now() - INTERVAL 1 DAY
        GROUP BY ALL
        ORDER BY bucket_start, country, event_type
        SETTINGS final = 1;  -- ou use SELECT ... FINAL
