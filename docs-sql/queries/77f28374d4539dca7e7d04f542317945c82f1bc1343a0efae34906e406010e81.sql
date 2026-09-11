SELECT COLUMNS('.*_amount|fee|tax')
FROM nyc_taxi.trips
ORDER BY rand() 
LIMIT 3;
