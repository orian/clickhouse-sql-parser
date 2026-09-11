SELECT rowNumberInBlock()
FROM
(
    SELECT *
    FROM system.numbers_mt
    LIMIT 10
) SETTINGS max_block_size = 2
