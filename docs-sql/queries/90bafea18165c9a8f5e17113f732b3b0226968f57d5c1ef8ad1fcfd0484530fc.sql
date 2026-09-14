CREATE TABLE hits_delta
    ENGINE = DeltaLake('https://datasets-documentation.s3.amazonaws.com/lake_formats/delta_lake/')

SELECT
    URL,
    count() AS cnt
FROM hits_delta
GROUP BY URL
ORDER BY cnt DESC
LIMIT 5
