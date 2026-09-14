
        SELECT count() OVER (PARTITION BY k) FROM (SELECT number % 8 AS k FROM numbers(50000000)) FORMAT Null;
    