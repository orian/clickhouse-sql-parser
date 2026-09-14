    taxi=# explain SELECT
            count(1) AS total,
            "Borough"
        FROM taxi.trips
        JOIN taxi.taxi_zone_dictionary
          ON trips.pickup_nyct2010_gid = toUInt64(taxi.taxi_zone_dictionary."LocationID")
        WHERE pickup_nyct2010_gid > 0
          AND dropoff_nyct2010_gid IN (132, 138)
        GROUP BY "Borough"
        ORDER BY total DESC;
                                  QUERY PLAN
    -----------------------------------------------------------------------
     Foreign Scan  (cost=1.00..5.10 rows=1000 width=40)
       Relations: Aggregate on ((trips) INNER JOIN (taxi_zone_dictionary))
    (2 rows)
    Time: 2.012 ms
