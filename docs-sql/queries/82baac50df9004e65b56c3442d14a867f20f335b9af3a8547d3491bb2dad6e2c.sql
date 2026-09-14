SELECT
    url,
    count() AS cnt
FROM icebergS3Cluster(
    'default',
    'https://datasets-documentation.s3.amazonaws.com/lake_formats/iceberg/'
)
GROUP BY url
ORDER BY cnt DESC
LIMIT 5
