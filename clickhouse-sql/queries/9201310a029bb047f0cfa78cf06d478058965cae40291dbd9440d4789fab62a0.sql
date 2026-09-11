
        INSERT INTO ts_grid_ingest
        SELECT
            number % 10000 AS id,
            toDateTime64(intDiv(number, 10000) * 15, 3) AS timestamp,
            (intDiv(number, 10000) % 1000)::Float64 AS value
        FROM numbers_mt(10000 * 5000)
    