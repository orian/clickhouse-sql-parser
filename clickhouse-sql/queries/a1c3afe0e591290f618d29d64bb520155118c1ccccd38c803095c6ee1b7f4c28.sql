
        INSERT INTO test_strings_{cardinality}
        SELECT lc AS col1, lc AS col2
        FROM
        (
            SELECT toLowCardinality(toString(number % {cardinality})) AS lc
            FROM numbers(10000000)
        )
    