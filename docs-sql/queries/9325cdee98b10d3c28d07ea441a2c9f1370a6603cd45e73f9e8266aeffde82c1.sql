   taxi=# SELECT *
   FROM taxi.trips
   JOIN taxi.taxi_zone_dictionary
       ON trips.dropoff_nyct2010_gid = taxi.taxi_zone_dictionary."LocationID"
   WHERE tip_amount > 0
   ORDER BY tip_amount DESC
   LIMIT 1000;
