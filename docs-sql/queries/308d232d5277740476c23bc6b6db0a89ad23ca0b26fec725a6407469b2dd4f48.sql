WITH
    10 AS length,
    (number / 1000000000.) * length AS x
SELECT pow((2 * length) * avg(exp(-(x * x))), 2) AS pi
FROM numbers_mt(1000000000.)

┌─────────────────pi─┐
│ 3.1415926890388595 │
└────────────────────┘

1 row in set. Elapsed: 1.245 sec. Processed 1.00 billion rows, 8.00 GB (803.25 million rows/s., 6.43 GB/s.)
