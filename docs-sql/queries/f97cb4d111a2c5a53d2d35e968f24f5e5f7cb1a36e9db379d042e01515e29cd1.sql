SELECT
    table_1.id,
    table_1.value,
    table_2.value
FROM table_1
LEFT JOIN table_2 ON table_1.id = table_2.id
WHERE table_2.id >= 2
ORDER BY table_1.id;
