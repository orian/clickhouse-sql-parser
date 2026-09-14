-- The data is processed in blocks of two rows: rowNumberInBlock restarts from 0 in every block,
-- while rowNumberInAllBlocks keeps counting across them.
SELECT
    number,
    rowNumberInBlock(),
    rowNumberInAllBlocks()
FROM system.numbers
LIMIT 6
SETTINGS max_block_size = 2
