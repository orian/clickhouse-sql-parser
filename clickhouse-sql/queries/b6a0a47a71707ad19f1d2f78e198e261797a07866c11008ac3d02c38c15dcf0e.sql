
        INSERT INTO ts_bucketing
        SELECT
            number % 2500 AS id,
            toDateTime64(intDiv(number, 2500) * 15, 3) AS timestamp,
            (intDiv(number, 2500) % 1000)::Float64 AS value
        FROM numbers_mt(2500 * 20000)
    