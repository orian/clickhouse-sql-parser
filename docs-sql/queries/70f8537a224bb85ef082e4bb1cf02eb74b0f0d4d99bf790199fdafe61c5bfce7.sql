CREATE TABLE nyc_taxi.trips_small_pk
(
    vendor_id LowCardinality(String),
    pickup_datetime DateTime('UTC'),
    dropoff_datetime DateTime('UTC'),
    passenger_count UInt8,
    trip_distance Float32,
    ratecode_id LowCardinality(Nullable(String)),
    pickup_location_id LowCardinality(String),
    dropoff_location_id LowCardinality(String),
    payment_type Nullable(UInt8),
    fare_amount Decimal32(2),
    extra Decimal32(2),
    mta_tax Nullable(Decimal32(2)),
    tip_amount Decimal32(2),
    tolls_amount Decimal32(2),
    total_amount Decimal32(2)
)
ENGINE = MergeTree
ORDER BY (passenger_count, pickup_datetime, dropoff_datetime);

INSERT INTO nyc_taxi.trips_small_pk
SELECT *
FROM nyc_taxi.trips_small_no_pk;
