  taxi=# SELECT
      pickup_ntaname,
      date_part('hour', pickup_datetime) as pickup_hour,
      SUM(1) AS pickups
  FROM taxi.trips
  WHERE pickup_ntaname != ''
  GROUP BY pickup_ntaname, pickup_hour
  ORDER BY pickup_ntaname, date_part('hour', pickup_datetime)
  LIMIT 5;
   pickup_ntaname | pickup_hour | pickups
  ----------------+-------------+---------
   Airport        |           0 |    3509
   Airport        |           1 |    1184
   Airport        |           2 |     401
   Airport        |           3 |     152
   Airport        |           4 |     213
  (5 rows)

  Time: 36.895 ms
