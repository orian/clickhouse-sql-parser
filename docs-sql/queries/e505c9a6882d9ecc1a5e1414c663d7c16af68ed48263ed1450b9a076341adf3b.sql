SELECT
    concat(database, '.', table),
    partition_id,
    count()
FROM system.parts
WHERE active
GROUP BY
    database,
    table,
    partition_id
