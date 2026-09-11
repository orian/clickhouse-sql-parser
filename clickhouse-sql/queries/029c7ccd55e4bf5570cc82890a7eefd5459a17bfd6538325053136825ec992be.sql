
        INSERT INTO matview_10000
        SELECT a, uniqState(b) b_count
        FROM
        (
            SELECT toString(intDiv(number, 20000)) a, toString(number % 10000) b
            FROM numbers_mt(20000000)
        )
        GROUP BY a
        SETTINGS max_insert_threads=8;
    