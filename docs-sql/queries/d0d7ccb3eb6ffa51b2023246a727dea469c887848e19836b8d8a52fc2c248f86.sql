system.parts テーブルを使用して、crypto_prices テーブルのパーツがどのディスクにあるかを確認します：

SELECT
    name,
    disk_name
FROM system.parts
WHERE (table = 'my_table') AND (active = 1)
