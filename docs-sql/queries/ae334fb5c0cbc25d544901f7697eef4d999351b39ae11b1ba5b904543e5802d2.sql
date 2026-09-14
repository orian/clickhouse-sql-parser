SELECT
    uniq(ratecode_id),
    uniq(pickup_location_id),
    uniq(dropoff_location_id),
    uniq(vendor_id)
FROM nyc_taxi.trips_small_inferred
FORMAT VERTICAL;
