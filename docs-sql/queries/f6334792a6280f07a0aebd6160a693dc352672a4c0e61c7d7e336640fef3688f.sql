Usando a tabela system.parts, veja em quais discos as partes estão para a tabela crypto_prices:

SELECT
    name,
    disk_name
FROM system.parts
WHERE (table = 'my_table') AND (active = 1)
