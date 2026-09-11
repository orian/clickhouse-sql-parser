Usando la tabla system.parts, consulte en qué discos se encuentran las partes de la tabla crypto_prices:

SELECT
    name,
    disk_name
FROM system.parts
WHERE (table = 'my_table') AND (active = 1)
