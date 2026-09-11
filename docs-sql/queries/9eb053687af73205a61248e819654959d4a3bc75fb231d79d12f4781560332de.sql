SELECT number AS key, 5 * number value, 'original' AS source
FROM numbers(16) WHERE key % 5 == 0
ORDER BY key WITH FILL;
