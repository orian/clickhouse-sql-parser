
        INSERT INTO distinct_high_cardinality_lc
        SELECT
            number AS id,
            toLowCardinality(toString(number % 5000)) AS lc
        FROM numbers(10000000);
    