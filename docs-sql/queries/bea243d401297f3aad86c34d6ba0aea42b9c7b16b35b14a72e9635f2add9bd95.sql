    SELECT
        table,
        count() AS parts,
        sum(rows) AS total_rows,
        formatReadableSize(sum(bytes_on_disk)) AS compressed_size
    FROM system.parts
    WHERE table = 'uk_price_paid'
      AND active = true
    GROUP BY table;
