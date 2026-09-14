
        INSERT INTO ts_bucketing_unsorted
        SELECT
            number % 2500 AS id,
            intHash64(number) AS h,
            toDateTime64(intDiv(number, 2500) * 15, 3) AS timestamp,
            (intDiv(number, 2500) % 1000)::Float64 AS value
        FROM numbers_mt(2500 * 8000)
    