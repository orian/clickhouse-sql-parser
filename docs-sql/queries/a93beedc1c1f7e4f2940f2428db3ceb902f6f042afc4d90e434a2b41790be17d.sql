-- Which countries have the most holidays?
SELECT countryOrRegion, count() AS holidays
FROM publicholidays
GROUP BY countryOrRegion
ORDER BY holidays DESC
LIMIT 10;
