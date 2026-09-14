
        INSERT INTO tab SELECT
            number as id,
            '',
            repeat('aaaaaaaaaaaaaaaa', 1),
            repeat('aaaaaaaaaaaaaaaa', 16),
            repeat('aaaaaaaaaaaaaaaa', 16 * 16),
            repeat('ääääääääääääääää', 1),
            repeat('ääääääääääääääää', 16),
            repeat('ääääääääääääääää', 16 * 16),
        FROM numbers(1000000)
    