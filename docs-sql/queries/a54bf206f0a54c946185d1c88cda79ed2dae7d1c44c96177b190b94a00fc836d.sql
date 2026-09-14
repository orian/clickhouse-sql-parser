WITH year(toStartOfInterval(date, toIntervalYear(10))) AS year
SELECT
    county,
    medianMap(map(year, price)) AS medianPrices
FROM uk.uk_price_paid
GROUP BY ALL
ORDER BY max(price) DESC
LIMIT 10;
