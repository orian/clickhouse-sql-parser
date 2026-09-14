SELECT
    table,
    column,
    part_type,
    sum(rows),
    sum(column_data_compressed_bytes),
    sum(column_data_uncompressed_bytes)
FROM system.parts_columns
WHERE active
GROUP BY
    table,
    column,
    part_type
ORDER BY
    table ASC,
    column ASC,
    part_type ASC
