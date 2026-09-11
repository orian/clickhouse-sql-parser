SELECT 
    COLUMNS('.*_amount'),
    COLUMNS('.*_date.*')
FROM nyc_taxi.trips
LIMIT 5;
