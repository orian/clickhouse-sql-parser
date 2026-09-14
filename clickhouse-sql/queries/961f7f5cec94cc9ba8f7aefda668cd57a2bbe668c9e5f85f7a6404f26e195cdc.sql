
        SELECT cume_dist() OVER (ORDER BY k) FROM (SELECT number % 2000 AS k FROM numbers(1000000)) FORMAT Null;
    