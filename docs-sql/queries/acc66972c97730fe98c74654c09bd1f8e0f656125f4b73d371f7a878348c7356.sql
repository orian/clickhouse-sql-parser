EXPLAIN indexes = 1
SELECT
    payment_type,
    count()
FROM nyc_taxi.trips_small_inferred
WHERE pickup_datetime >= '2009-01-01'
  AND pickup_datetime < '2009-04-01'
GROUP BY payment_type
SETTINGS
    use_query_condition_cache = 0,
    use_skip_indexes_on_data_read = 0;
