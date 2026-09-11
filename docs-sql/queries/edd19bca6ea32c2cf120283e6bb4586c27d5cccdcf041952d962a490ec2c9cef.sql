SELECT quantilesExact(0, 0.5, 0.75, 0.99)(number) AS quants
FROM foo.randomNumbers
