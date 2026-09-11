
        SELECT cramersV(number, number) OVER (ORDER BY number) FROM numbers(5_000) FORMAT NULL;
    