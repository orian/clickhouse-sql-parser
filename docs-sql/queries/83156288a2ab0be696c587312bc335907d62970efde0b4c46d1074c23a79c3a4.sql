SELECT 4 * sum(if(number % 2, -1, 1) / ((number * 2) + 1)) AS pi
FROM numbers_mt(1000000000.)

┌────────────────pi─┐
│ 3.141592652589797 │
└───────────────────┘

1 row in set. Elapsed: 0.432 sec. Processed 1.00 billion rows, 8.00 GB (2.32 billion rows/s., 18.53 GB/s.)
