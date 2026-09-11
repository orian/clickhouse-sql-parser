SELECT
    payment_type,
    count() AS trip_count,
    formatReadableQuantity(sum(trip_distance)) AS total_distance,
    avg(total_amount) AS total_amount_avg,
    avg(tip_amount) AS tip_amount_avg
FROM nyc_taxi.trips_small_inferred
WHERE pickup_datetime >= '2009-01-01'
  AND pickup_datetime < '2009-04-01'
GROUP BY payment_type
ORDER BY trip_count DESC;
