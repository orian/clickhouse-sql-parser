
        CREATE TABLE agg_deserialize
        (
            t DateTime,
            v1 AggregateFunction(avgState, UInt64),
            v2 AggregateFunction(argMax, UInt64, DateTime)
        )
        ENGINE = MergeTree() ORDER BY t
    