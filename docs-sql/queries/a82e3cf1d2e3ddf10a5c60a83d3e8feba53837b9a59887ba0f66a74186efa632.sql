SELECT
    *
FROM file(
    'squares.buffers',
    'Buffers',
    'col_1 UInt64, col_2 UInt64'
);
