SELECT
    part_name,
    max(mark_number) AS entries
FROM mergeTreeIndex('uk', 'uk_price_paid_simple')
GROUP BY part_name;
