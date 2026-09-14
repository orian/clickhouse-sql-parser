使用 system.parts 表，查看 crypto_prices 表中各 parts 所在的磁盘：

SELECT
    name,
    disk_name
FROM system.parts
WHERE (table = 'my_table') AND (active = 1)
