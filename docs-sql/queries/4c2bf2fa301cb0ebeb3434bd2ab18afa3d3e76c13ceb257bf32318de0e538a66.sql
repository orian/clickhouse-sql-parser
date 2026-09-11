SELECT
    count(*),
    sum(data_uncompressed_bytes)
FROM system.snapshot_parts
WHERE (name, table_id) NOT IN (
    SELECT
        name,
        toString(tables.uuid)
    FROM system.parts
    INNER JOIN system.tables ON (parts.`table` = tables.name) AND parts.active
)
