CREATE TABLE nyc_taxi.trips_for_payment_dashboard
ENGINE = MergeTree
ORDER BY (payment_type, pickup_datetime)
AS SELECT
    assumeNotNull(source.payment_type) AS payment_type,
    assumeNotNull(source.pickup_datetime) AS pickup_datetime,
    trip_distance,
    total_amount
FROM nyc_taxi.trips_small_inferred AS source
WHERE source.payment_type IS NOT NULL
  AND source.pickup_datetime IS NOT NULL;
