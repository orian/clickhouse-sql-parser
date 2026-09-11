SELECT
    partition,
    count() AS parts,
    sum(rows) AS rows
FROM system.parts
WHERE (database = 'uk') AND (`table` = 'uk_price_paid_simple_partitioned') AND active
GROUP BY partition
ORDER BY partition ASC;
