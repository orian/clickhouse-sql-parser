
        CREATE TABLE ts_deserialization
        (
            id UInt16,
            state AggregateFunction(timeSeriesRateToGrid(0, 10000, 100, 300), DateTime, Float64)
        )
        ENGINE = MergeTree
        ORDER BY id
    