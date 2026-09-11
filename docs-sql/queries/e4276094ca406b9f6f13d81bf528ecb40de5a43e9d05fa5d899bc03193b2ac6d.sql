WITH year(toStartOfInterval(date, toIntervalYear(10))) AS year
SELECT
    county,
    medianMap(map(year, price)) AS medianPrices,
    mapApply((k, v) -> (k, floor(v)), avgMap(map(year, price))) AS avgPrices
FROM uk.uk_price_paid
WHERE year >= 2010
GROUP BY ALL
ORDER BY max(price) DESC
LIMIT 10;
