WITH toDate('2018-01-01') AS date SELECT subtractTupleOfIntervals(date, (INTERVAL 1 DAY, INTERVAL 1 YEAR))
