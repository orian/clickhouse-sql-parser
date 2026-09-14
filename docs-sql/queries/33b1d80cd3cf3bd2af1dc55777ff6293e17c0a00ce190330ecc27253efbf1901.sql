    INSERT INTO nyc_taxi.trips_small
    SELECT
        trip_id,
        pickup_datetime,
        dropoff_datetime,
        pickup_longitude,
        pickup_latitude,
        dropoff_longitude,
        dropoff_latitude,
        passenger_count,
        trip_distance,
        fare_amount,
        extra,
        tip_amount,
        tolls_amount,
        total_amount,
        payment_type,
        pickup_ntaname,
        dropoff_ntaname
    FROM s3(
        'https://datasets-documentation.s3.eu-west-3.amazonaws.com/nyc-taxi/trips_{0..2}.gz',
        'TabSeparatedWithNames'
    );
