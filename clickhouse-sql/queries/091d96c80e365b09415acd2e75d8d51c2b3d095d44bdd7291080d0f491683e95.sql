
        SELECT cume_dist() OVER (ORDER BY k) FROM (SELECT number % 5000 AS k FROM numbers(1000000)) FORMAT Null;
    