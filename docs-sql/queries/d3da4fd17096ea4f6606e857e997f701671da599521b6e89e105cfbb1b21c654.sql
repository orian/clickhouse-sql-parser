SELECT 3 + (4 * sum(if((number % 2) = 0, if((number % 4) = 0, -1 / ((number * (number + 1)) * (number + 2)), 1 / ((number * (number + 1)) * (number + 2))), 0))) AS pi
FROM numbers_mt(2, 10000000000)

┌─────────────────pi─┐
│ 3.1415926525808087 │
└────────────────────┘

1 row in set. Elapsed: 9.825 sec. Processed 10.00 billion rows, 80.00 GB (1.02 billion rows/s., 8.14 GB/s.)
