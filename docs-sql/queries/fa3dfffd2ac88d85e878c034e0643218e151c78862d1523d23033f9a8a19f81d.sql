INSERT INTO logs SELECT
 ['200', '404', '502', '403'][toInt32(randBinomial(4, 0.1)) + 1] AS code,
    now() + toIntervalMinute(number) AS timestamp
FROM numbers(100000000)

0 rows in set. Elapsed: 15.845 sec. Processed 100.00 million rows, 800.00 MB (6.31 million rows/s., 50.49 MB/s.)

SELECT count()
FROM logs

┌───count()─┐
│ 100000000 │ -- 1 亿
└───────────┘

1 row in set. Elapsed: 0.002 sec.
