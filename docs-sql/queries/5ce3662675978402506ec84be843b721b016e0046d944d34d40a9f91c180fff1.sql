    taxi=# SELECT
        count(1) AS total,
        "Borough"
    FROM taxi.trips
    JOIN taxi.taxi_zone_dictionary
      ON trips.pickup_nyct2010_gid = toUInt64(taxi.taxi_zone_dictionary."LocationID")
    WHERE pickup_nyct2010_gid > 0
      AND dropoff_nyct2010_gid IN (132, 138)
    GROUP BY "Borough"
    ORDER BY total DESC;
     total | borough_name
    -------+---------------
      7053 | Manhattan
      6828 | Brooklyn
      4458 | Queens
      2670 | Bronx
       554 | Staten Island
        53 | EWR
    (6 rows)

    Time: 48.449 ms
