        SELECT
        partition,
        name,
        rows,
        bytes_on_disk,
        marks
        FROM system.parts
        WHERE table = 'uk_price_paid'
        AND active = true
        ORDER BY partition
        LIMIT 20;
