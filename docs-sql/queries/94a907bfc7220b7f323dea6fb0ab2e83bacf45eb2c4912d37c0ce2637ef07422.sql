FROM nyc_taxi.trips
SELECT COLUMNS('.*_amount')
LIMIT 10;
