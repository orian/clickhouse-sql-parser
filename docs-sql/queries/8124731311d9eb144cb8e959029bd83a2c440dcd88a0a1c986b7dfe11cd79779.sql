-- Input table with temperature data
SELECT exponentialMovingAverage(5)(temperature, timestamp)
FROM VALUES('temperature Int32, timestamp Int32',
    (95, 1), (95, 2), (95, 3), (96, 4), (96, 5), (96, 6), (96, 7),
    (97, 8), (97, 9), (97, 10), (97, 11), (98, 12), (98, 13), (98, 14),
    (98, 15), (99, 16), (99, 17), (99, 18), (100, 19), (100, 20))
