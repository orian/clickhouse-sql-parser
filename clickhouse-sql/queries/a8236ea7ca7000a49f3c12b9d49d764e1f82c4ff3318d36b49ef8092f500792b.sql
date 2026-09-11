
        INSERT INTO ts_scalar_add
        SELECT
            (number % 200)::UInt16 AS id,
            toDateTime64(1000000 + intDiv(number, 200) * 15, 3) AS timestamp,
            (number % 1000)::Float64 AS value
        FROM numbers_mt(40000000)
    