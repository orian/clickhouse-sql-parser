
        SELECT number, avg(DISTINCT number) OVER () FROM numbers(0, 100_000) FORMAT Null;
    