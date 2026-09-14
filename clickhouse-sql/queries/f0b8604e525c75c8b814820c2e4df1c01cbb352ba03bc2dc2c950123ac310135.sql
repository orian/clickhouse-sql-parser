
        INSERT INTO ts_grid_sparse
        SELECT
            (number % 400)::UInt16 AS id,
            toDateTime(intDiv(number, 400) * 100) AS timestamp,
            (intDiv(number, 400) % 1000)::Float64 AS value
        FROM numbers_mt(400 * 1000)
    