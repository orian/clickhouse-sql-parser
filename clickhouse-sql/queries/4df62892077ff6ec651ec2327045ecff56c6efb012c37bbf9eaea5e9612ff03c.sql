
        INSERT INTO aov_states
        SELECT
            intDiv(number, 100000) % 200 AS key,
            analysisOfVarianceState(number * 1.5, (number % 100000)::UInt32)
        FROM numbers_mt(20000000)
        GROUP BY key
    