        SELECT
        county,
        round(avg(price)) AS avg_price,
        count()           AS sales
        FROM uk_price_paid
        GROUP BY county
        ORDER BY avg_price DESC;
