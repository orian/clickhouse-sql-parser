    taxi=# SELECT
        avg(tip_amount) AS avg_tip,
        avg(fare_amount) AS avg_fare,
        avg(passenger_count) AS avg_passenger,
        count(*) AS count,
        round((date_part('epoch', dropoff_datetime) - date_part('epoch', pickup_datetime)) / 60) as trip_minutes
    FROM taxi.trips
    WHERE round((date_part('epoch', dropoff_datetime) - date_part('epoch', pickup_datetime)) / 60) > 0
    GROUP BY trip_minutes
    ORDER BY trip_minutes DESC
    LIMIT 5;
          avg_tip      |     avg_fare     |  avg_passenger   | count | trip_minutes
    -------------------+------------------+------------------+-------+--------------
                  1.96 |                8 |                1 |     1 |        27512
                     0 |               12 |                2 |     1 |        27500
     0.562727272727273 | 17.4545454545455 | 2.45454545454545 |    11 |         1440
     0.716564885496183 | 14.2786259541985 | 1.94656488549618 |   131 |         1439
      1.00945205479452 | 12.8787671232877 | 1.98630136986301 |   146 |         1438
    (5 rows)

    Time: 45.477 ms
