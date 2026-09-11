
        INSERT INTO str_sort_long
        SELECT
            randomString(64) AS a,
            rand() AS b, cast(randomString(128) AS FixedString(128)) as c,
            randomString(256) as d,
            rand() % 1000 as e,
            rand() % 1000 as f
        FROM numbers(5e6)
    