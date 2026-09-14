SELECT DestCityName, uniqExact(OriginCityName) AS u
FROM ontime
WHERE Year >= 2000 AND Year <= 2010
GROUP BY DestCityName
ORDER BY u DESC LIMIT 10;
