SELECT
    county,
    avg(price)
FROM uk.uk_price_paid_with_projections
GROUP BY county
ORDER BY avg(price) DESC
LIMIT 3
