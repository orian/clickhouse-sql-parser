SELECT
    URL,
    count() AS cnt
FROM deltaLakeCluster(
    'default',
    'https://datasets-documentation.s3.amazonaws.com/lake_formats/delta_lake/'
)
GROUP BY URL
ORDER BY cnt DESC
LIMIT 5
