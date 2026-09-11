    CREATE TABLE hits_iceberg
        ENGINE = IcebergS3('https://datasets-documentation.s3.amazonaws.com/lake_formats/iceberg/')

    SELECT
        url,
        count() AS cnt
    FROM hits_iceberg
    GROUP BY url
    ORDER BY cnt DESC
    LIMIT 5
