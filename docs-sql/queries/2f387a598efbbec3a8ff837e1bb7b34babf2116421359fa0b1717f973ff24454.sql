  taxi=# SELECT
      pickup_date,
      pickup_ntaname,
      SUM(1) AS number_of_trips
  FROM taxi.trips
  GROUP BY pickup_date, pickup_ntaname
  ORDER BY pickup_date ASC LIMIT 10;
   pickup_date |         pickup_ntaname         | number_of_trips
  -------------+--------------------------------+-----------------
   2015-07-01  | Williamsburg                   |               1
   2015-07-01  | park-cemetery-etc-Queens       |               6
   2015-07-01  | Maspeth                        |               1
   2015-07-01  | Stuyvesant Town-Cooper Village |              44
   2015-07-01  | Rego Park                      |               1
   2015-07-01  | Greenpoint                     |               7
   2015-07-01  | Highbridge                     |               1
   2015-07-01  | Briarwood-Jamaica Hills        |               3
   2015-07-01  | Airport                        |             550
   2015-07-01  | East Harlem North              |              32
  (10 rows)

  Time: 30.978 ms
