
        INSERT INTO ts_deserialization
        SELECT
            (number % 1000)::UInt16 AS id,
            timeSeriesRateToGridState(0, 10000, 100, 300)(
                toDateTime(intDiv(number, 1000)),
                (number % 1000)::Float64) AS state
        FROM numbers(1000 * 10000)
        GROUP BY id
    