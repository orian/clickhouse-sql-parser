SELECT count() AS c
FROM logs
WHERE code = '200'

┌────────c─┐
│ 65607542 │ -- 65.61 million
└──────────┘

1 row in set. Elapsed: 0.021 sec. Processed 49.15 thousand rows, 49.17 KB (2.34 million rows/s., 2.34 MB/s.)
Peak memory usage: 92.70 KiB.
