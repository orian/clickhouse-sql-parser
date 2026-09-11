  taxi=# SELECT
          passenger_count,
          avg(total_amount)::NUMERIC(10, 2) AS average_total_amount
      FROM taxi.trips
      GROUP BY passenger_count;
   passenger_count | average_total_amount
  -----------------+----------------------
                 0 |                22.68
                 1 |                15.96
                 2 |                17.14
                 3 |                16.75
                 4 |                17.32
                 5 |                16.34
                 6 |                16.03
                 7 |                59.79
                 8 |                36.40
                 9 |                 9.79
  (10 rows)

  Time: 27.266 ms
