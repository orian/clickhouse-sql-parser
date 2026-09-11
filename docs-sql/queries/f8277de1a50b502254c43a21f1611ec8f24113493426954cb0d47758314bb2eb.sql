SELECT
    min(payment_type),
    max(payment_type),
    min(passenger_count),
    max(passenger_count)
FROM nyc_taxi.trips_small_inferred;
