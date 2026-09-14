
        SELECT number, avg(DISTINCT number) OVER (ORDER BY number) FROM numbers(0, 100_000) FORMAT Null;
    