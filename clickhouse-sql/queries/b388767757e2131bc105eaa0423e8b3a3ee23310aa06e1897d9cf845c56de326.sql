
        CREATE TABLE test(
            t UInt64,
            q1 AggregateFunction(quantilesTiming(0.50, 0.75, 0.90, 0.99), Float64),
            q2 AggregateFunction(quantilesTiming(0.50, 0.75, 0.90, 0.99), Float64),
            q3 AggregateFunction(quantilesTiming(0.50, 0.75, 0.90, 0.99), Float64),
            q4 AggregateFunction(quantilesTiming(0.50, 0.75, 0.90, 0.99), Float64),
            q5 AggregateFunction(quantilesTiming(0.50, 0.75, 0.90, 0.99), Float64)
        ) ENGINE=SummingMergeTree()
        ORDER BY t
    