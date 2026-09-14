WITH
    PER_MINUTE AS
    (
    SELECT
        toStartOfInterval(modification_time, toIntervalMinute(1)) AS t,
        count() AS new_part_count
    FROM
        clusterAllReplicas(default, merge(system, '^parts'))
    WHERE
        (database = 'default') AND
        (table = 'your_table') AND
        (active = true) AND
        (level = 0) AND
        (modification_time >= (now() - toIntervalHour(1)))
    GROUP BY
        t
    ORDER BY
        t ASC
    SETTINGS skip_unavailable_shards = 1
    )
SELECT
    AVG(new_part_count) AS new_parts_per_minute,
    new_parts_per_minute / 60 AS new_parts_per_second
FROM
    PER_MINUTE
