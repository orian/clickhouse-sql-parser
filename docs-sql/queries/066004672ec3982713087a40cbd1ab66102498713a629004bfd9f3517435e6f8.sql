SELECT
    number AS num,
    number * number AS num_square
FROM numbers(10)
INTO OUTFILE 'squares.buffers'
FORMAT Buffers;
