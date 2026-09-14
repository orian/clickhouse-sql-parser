CREATE TABLE nyc_taxi.trips_by_payment_type
(
    payment_type Int64,
    trip_count UInt64
)
ENGINE = MergeTree
ORDER BY payment_type;

CREATE MATERIALIZED VIEW nyc_taxi.trips_by_payment_type_mv
REFRESH EVERY 1 HOUR
TO nyc_taxi.trips_by_payment_type
AS SELECT
    assumeNotNull(source.payment_type) AS payment_type,
    count() AS trip_count
FROM nyc_taxi.trips_small_inferred AS source
WHERE source.payment_type IS NOT NULL
GROUP BY payment_type;
