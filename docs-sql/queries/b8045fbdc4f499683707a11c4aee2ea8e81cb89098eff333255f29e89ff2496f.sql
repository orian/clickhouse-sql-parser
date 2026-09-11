%%sql --save short-trips --no-execute
SELECT *
FROM trips
WHERE trip_distance < 6.3
