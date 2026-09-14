ALTER TABLE nyc_taxi.trips_small_inferred
ADD PROJECTION trips_by_payment_type
(
    SELECT
        payment_type,
        pickup_datetime,
        trip_distance,
        total_amount
    ORDER BY (payment_type, pickup_datetime)
);

ALTER TABLE nyc_taxi.trips_small_inferred
MATERIALIZE PROJECTION trips_by_payment_type;
