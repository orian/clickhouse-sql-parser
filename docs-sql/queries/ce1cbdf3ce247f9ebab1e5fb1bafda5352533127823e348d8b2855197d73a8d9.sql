    SELECT
        passenger_count,
        ceil(avg(total_amount),2) AS average_total_amount
    FROM trips
    GROUP BY passenger_count
