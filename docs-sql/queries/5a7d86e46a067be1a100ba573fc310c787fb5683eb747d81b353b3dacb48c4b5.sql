SELECT passenger_count, avg(total_amount) FROM trips_mergetree GROUP BY passenger_count;
