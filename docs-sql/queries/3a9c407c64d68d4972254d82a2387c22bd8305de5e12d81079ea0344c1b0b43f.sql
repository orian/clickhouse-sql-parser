SELECT
    database,
    table,
    partition_id,
    count() AS active_parts,
    sum(rows) AS rows,
    formatReadableSize(sum(bytes_on_disk)) AS size_on_disk
FROM system.parts
WHERE active
  AND database = '<database_name>'
  AND table = '<table_name>'
GROUP BY
    database,
    table,
    partition_id
ORDER BY active_parts DESC;
