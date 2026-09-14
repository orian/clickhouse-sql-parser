        SELECT
        toYear(date) AS year,
        round(avg(price)) AS avg_price,
        count() AS sales,
        min(price) AS cheapest,
        max(price) AS most_expensive
        FROM uk_price_paid
        WHERE postcode LIKE 'SW1A%'
        GROUP BY year
        ORDER BY year DESC;
