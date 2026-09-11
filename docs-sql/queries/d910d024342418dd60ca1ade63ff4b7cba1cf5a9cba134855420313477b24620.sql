SELECT
    table,
    part_type,
    count(*)
FROM system.parts
WHERE active
GROUP BY
    table,
    part_type
ORDER BY
    table ASC,
    part_type ASC
