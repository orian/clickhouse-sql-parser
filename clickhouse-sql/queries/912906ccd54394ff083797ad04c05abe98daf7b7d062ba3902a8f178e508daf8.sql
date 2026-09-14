
        INSERT INTO test
        SELECT
            number / 10 as t,
            quantilesTimingState(0.50, 0.75, 0.90, 0.99)(number/1000) as q1,
            quantilesTimingState(0.50, 0.75, 0.90, 0.99)(number/1000) as q2,
            quantilesTimingState(0.50, 0.75, 0.90, 0.99)(number/1000) as q3,
            quantilesTimingState(0.50, 0.75, 0.90, 0.99)(number/1000) as q4,
            quantilesTimingState(0.50, 0.75, 0.90, 0.99)(number/1000) as q5
        FROM numbers(1000 * 1000)
        GROUP BY t
    