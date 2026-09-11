
        CREATE TABLE aov_states
        (
            key UInt32,
            state AggregateFunction(analysisOfVariance, Float64, UInt32)
        )
        ENGINE = MergeTree ORDER BY key
    