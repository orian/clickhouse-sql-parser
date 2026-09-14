   taxi=# SELECT
           count(1) AS total,
           COALESCE(NULLIF(dictGet(
               'taxi.taxi_zone_dictionary', 'Borough',
               toUInt64(pickup_nyct2010_gid)
           ), ''), 'Unknown') AS borough_name
       FROM taxi.trips
       WHERE dropoff_nyct2010_gid = 132 OR dropoff_nyct2010_gid = 138
       GROUP BY borough_name
       ORDER BY total DESC;
    total | borough_name
   -------+---------------
    23683 | Unknown
     7053 | Manhattan
     6828 | Brooklyn
     4458 | Queens
     2670 | Bronx
      554 | Staten Island
       53 | EWR
   (7 rows)

   Time: 66.245 ms
