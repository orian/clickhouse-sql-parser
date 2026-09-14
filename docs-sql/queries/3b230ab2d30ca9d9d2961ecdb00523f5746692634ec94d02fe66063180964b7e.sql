SELECT
    passenger_count,
    avg(trip_distance) AS distance,
    count() AS c
FROM nyc_taxi.trips_small
GROUP BY passenger_count
ORDER BY passenger_count ASC
