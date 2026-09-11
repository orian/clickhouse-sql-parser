    SELECT
        url,
        count() AS cnt
    FROM hits_iceberg
    GROUP BY url
    ORDER BY cnt DESC
    LIMIT 5
