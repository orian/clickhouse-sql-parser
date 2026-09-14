  SELECT
    tip_amount,
    trip_id,
    dateDiff('minutes', pickup_datetime, dropoff_datetime) AS trip_duration_min
  FROM nyc_taxi.trips_with_projection WHERE tip_amount > 200 AND trip_duration_min > 0
  ORDER BY tip_amount, trip_id ASC
