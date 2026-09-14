SELECT
    pickup_datetime,
    payment_type,
    total_amount
FROM nyc_taxi.trips_small_inferred
WHERE pickup_datetime >= '2009-01-01'
  AND pickup_datetime < '2009-04-01'
LIMIT 1000;
