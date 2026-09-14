        SELECT
            toYear(date) AS year,
            round(avg(price)) AS avg_price,
            count() AS sales
        FROM uk_price_paid_by_town
        WHERE town = 'LONDON'
        GROUP BY year
        ORDER BY year DESC;
