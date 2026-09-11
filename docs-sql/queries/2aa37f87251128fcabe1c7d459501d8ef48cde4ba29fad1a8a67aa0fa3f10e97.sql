SELECT
    town,
    count() AS sales,
    round(avg(price)) AS average_price
FROM uk_price_paid
GROUP BY town
ORDER BY sales DESC
LIMIT 10
