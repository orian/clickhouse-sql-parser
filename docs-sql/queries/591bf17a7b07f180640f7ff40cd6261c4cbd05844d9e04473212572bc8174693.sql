SELECT *
FROM system.settings
WHERE name LIKE '%min_insert_block_size_%'
FORMAT Vertical
