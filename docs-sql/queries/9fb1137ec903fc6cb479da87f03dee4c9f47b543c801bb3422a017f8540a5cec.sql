EXPLAIN indexes = 1
SELECT
    payment_type,
    count() AS trip_count,
    formatReadableQuantity(sum(trip_distance)) AS total_distance,
    avg(total_amount) AS total_amount_avg,
    avg(tip_amount) AS tip_amount_avg
FROM nyc_taxi.trips_small_pk
WHERE pickup_datetime >= '2009-01-01'
  AND pickup_datetime < '2009-04-01'
GROUP BY payment_type
ORDER BY trip_count DESC
SETTINGS
    use_query_condition_cache = 0,
    use_skip_indexes_on_data_read = 0;
