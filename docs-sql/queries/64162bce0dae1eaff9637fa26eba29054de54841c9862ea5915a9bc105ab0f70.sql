Используя таблицу system.parts, проверьте, на каких дисках находятся части таблицы crypto_prices:

SELECT
    name,
    disk_name
FROM system.parts
WHERE (table = 'my_table') AND (active = 1)
