SELECT
   passenger_count,
   avg(total_amount)
FROM nyc_taxi.trips_small
WHERE passenger_count < 10
GROUP BY passenger_count;
