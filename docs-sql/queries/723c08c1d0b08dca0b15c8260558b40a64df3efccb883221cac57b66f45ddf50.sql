DROP TABLE drop table nyc_taxi.trips_small_dist;
CREATE VIEW nyc_taxi.trips_small_dist AS SELECT * FROM nyc_taxi.trips_small_dist_local;
SELECT count() from nyc_taxi.trips_small_dist;
3000317
