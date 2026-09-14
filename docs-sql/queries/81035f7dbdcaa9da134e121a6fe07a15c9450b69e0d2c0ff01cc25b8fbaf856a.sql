    taxi=# SET timezone = 'America/New_York';
    SET
    taxi=# SELECT
        pickup_datetime,
        dropoff_datetime,
        total_amount,
        pickup_nyct2010_gid,
        dropoff_nyct2010_gid,
        CASE
            WHEN dropoff_nyct2010_gid = 138 THEN 'LGA'
            WHEN dropoff_nyct2010_gid = 132 THEN 'JFK'
        END AS airport_code,
        EXTRACT(YEAR FROM pickup_datetime) AS year,
        EXTRACT(DAY FROM pickup_datetime) AS day,
        EXTRACT(HOUR FROM pickup_datetime) AS hour
    FROM taxi.trips
    WHERE dropoff_nyct2010_gid IN (132, 138)
    ORDER BY pickup_datetime
    LIMIT 5;
        pickup_datetime     |    dropoff_datetime    | total_amount | pickup_nyct2010_gid | dropoff_nyct2010_gid | airport_code | year | day | hour
    ------------------------+------------------------+--------------+---------------------+----------------------+--------------+------+-----+------
     2015-06-30 20:04:14-04 | 2015-06-30 20:15:29-04 |        13.30 |                 -34 |                  132 | JFK          | 2015 |  30 |   20
     2015-06-30 20:09:42-04 | 2015-06-30 20:12:55-04 |         6.80 |                  50 |                  138 | LGA          | 2015 |  30 |   20
     2015-06-30 20:23:04-04 | 2015-06-30 20:24:39-04 |         4.80 |                -125 |                  132 | JFK          | 2015 |  30 |   20
     2015-06-30 20:27:51-04 | 2015-06-30 20:39:02-04 |        14.72 |                -101 |                  138 | LGA          | 2015 |  30 |   20
     2015-06-30 20:32:03-04 | 2015-06-30 20:55:39-04 |        39.34 |                  48 |                  138 | LGA          | 2015 |  30 |   20
    (5 rows)

    Time: 17.450 ms
