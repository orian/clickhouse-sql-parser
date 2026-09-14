SELECT
    database,
    table,
    partition,
    count() AS parts,
    formatReadableSize(sum(bytes_on_disk)) AS bytes_on_disk,
    formatReadableQuantity(sum(rows)) AS rows,
    sum(marks) AS marks
FROM system.parts
WHERE (database != 'system') AND active
GROUP BY
    database,
    table,
    partition
ORDER BY database ASC
