
        INSERT INTO matview_1
        SELECT '1', uniqState(number) b_count
        FROM
        (
            SELECT *
            FROM numbers_mt(2000000)
        )
        GROUP BY number
        SETTINGS max_insert_threads=8;
    