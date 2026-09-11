SELECT 
    mark_number + 1 AS entry,
    town,
    street
FROM mergeTreeIndex('uk', 'uk_price_paid_simple')
WHERE part_name = (SELECT any(part_name) FROM mergeTreeIndex('uk', 'uk_price_paid_simple')) 
ORDER BY mark_number ASC
LIMIT 10;
