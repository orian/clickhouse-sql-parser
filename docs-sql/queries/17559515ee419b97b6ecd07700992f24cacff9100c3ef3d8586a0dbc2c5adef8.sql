WITH year(toStartOfInterval(date, toIntervalYear(10))) AS year
SELECT
    year,
    medianMap(map(postcode1 || ' ' || postcode2, price)) AS medianPrices
FROM uk.uk_price_paid
WHERE postcode1 LIKE 'NP1'
GROUP BY ALL;
