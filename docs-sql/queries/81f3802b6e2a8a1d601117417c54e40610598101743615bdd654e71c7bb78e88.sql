    SELECT
        name,
        count() AS parts,
        sum(rows) AS total_rows,
        formatReadableSize(sum(bytes_on_disk)) AS size
    FROM system.projection_parts
    WHERE table = 'uk_price_paid'
      AND active = true
    GROUP BY name;
