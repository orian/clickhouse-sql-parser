    -- Table function
    SELECT url, count() AS cnt
    FROM icebergS3('https://datasets-documentation.s3.amazonaws.com/lake_formats/iceberg/')
    GROUP BY url ORDER BY cnt DESC LIMIT 5

    -- Table engine
    SELECT url, count() AS cnt
    FROM hits_iceberg
    GROUP BY url ORDER BY cnt DESC LIMIT 5

    -- Catalog
    SELECT url, count() AS cnt
    FROM my_lake.`<database>.<table>`
    GROUP BY url ORDER BY cnt DESC LIMIT 5
