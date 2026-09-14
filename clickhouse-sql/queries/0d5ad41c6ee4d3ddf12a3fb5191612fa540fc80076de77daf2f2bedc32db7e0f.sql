
        SELECT number, s, min(DISTINCT s) OVER(PARTITION BY number % 5_000) AS min_varlen_string
        FROM
        (
            SELECT number, toString(number % 10_000) AS s
            FROM numbers(100_000)
        ) FORMAT Null;
    