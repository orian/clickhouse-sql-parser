      SELECT
          avg(tip_amount) AS avg_tip,
          avg(fare_amount) AS avg_fare,
          avg(passenger_count) AS avg_passenger,
          count() AS count,
          truncate(date_diff('second', pickup_datetime, dropoff_datetime)/60) as trip_minutes
      FROM trips
      WHERE trip_minutes > 0
      GROUP BY trip_minutes
      ORDER BY trip_minutes DESC
