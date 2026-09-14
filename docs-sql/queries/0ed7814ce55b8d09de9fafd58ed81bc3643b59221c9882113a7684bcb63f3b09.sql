       SELECT
           count(1) AS total,
           dictGetOrDefault('taxi_zone_dictionary','Borough', toUInt64(pickup_nyct2010_gid), 'Unknown') AS borough_name
       FROM trips
       WHERE dropoff_nyct2010_gid = 132 OR dropoff_nyct2010_gid = 138
       GROUP BY borough_name
       ORDER BY total DESC
