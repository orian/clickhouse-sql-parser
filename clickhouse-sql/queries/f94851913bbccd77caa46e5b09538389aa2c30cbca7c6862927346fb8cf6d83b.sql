
        INSERT INTO test_limit_by_{cardinality}
        SELECT number % {cardinality}, number % 97, number
        FROM numbers(5000000)
    