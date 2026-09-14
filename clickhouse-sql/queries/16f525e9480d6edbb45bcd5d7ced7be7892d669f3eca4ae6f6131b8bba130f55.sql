
        SELECT s, ts, argMax(DISTINCT ts, s) OVER(PARTITION BY ts % 5_000) AS value_with_max_ts
        FROM
        (
            SELECT number AS ts, toString(number % 10_000) AS s
            FROM numbers(100_000)
        ) FORMAT Null;
    