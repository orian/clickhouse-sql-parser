
        INSERT INTO agg_deserialize SELECT
            now() + number AS t,
            initializeAggregation('avgState', number),
            initializeAggregation('argMaxState', number, t)
        FROM numbers(50000000)
    