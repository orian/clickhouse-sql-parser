    EXPLAIN actions = 1, compact = 1, pretty = 1, indexes = 1
    WITH
        dateDiff('s', pickup_datetime, dropoff_datetime) AS trip_time,
        (trip_distance / trip_time) * 3600 AS speed_mph
    SELECT quantiles(0.5, 0.75, 0.9, 0.99)(trip_distance)
    FROM nyc_taxi.trips_small_inferred
    WHERE speed_mph > 30
