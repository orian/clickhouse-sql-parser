CREATE TABLE nyc_taxi.trips_by_day
(
    pickup_date Date,
    trip_count UInt64
)
ENGINE = SummingMergeTree
ORDER BY pickup_date;

CREATE MATERIALIZED VIEW nyc_taxi.trips_by_day_mv
TO nyc_taxi.trips_by_day
AS SELECT
    toDate(assumeNotNull(pickup_datetime)) AS pickup_date,
    count() AS trip_count
FROM nyc_taxi.trips_small_inferred
WHERE pickup_datetime IS NOT NULL
GROUP BY pickup_date;
