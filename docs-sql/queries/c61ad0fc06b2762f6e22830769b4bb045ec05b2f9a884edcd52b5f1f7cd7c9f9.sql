WITH random_points AS
    (
        SELECT (rand64(1) / pow(2, 64), rand64(2) / pow(2, 64)) AS point
        FROM numbers(1000000000)
    )
SELECT (4 * countIf(L2Norm(point) < 1)) / count() AS pi
FROM random_points

┌──────────pi─┐
│ 3.141627208 │
└─────────────┘

1 row in set. Elapsed: 4.742 sec. Processed 1.00 billion rows, 8.00 GB (210.88 million rows/s., 1.69 GB/s.)
