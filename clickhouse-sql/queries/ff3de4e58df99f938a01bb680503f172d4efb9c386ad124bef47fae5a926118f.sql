
        INSERT INTO distinct_low_cardinality_lc
        SELECT
            number AS id,
            toLowCardinality(toString(number % 100)) AS lc
        FROM numbers(10000000);
    