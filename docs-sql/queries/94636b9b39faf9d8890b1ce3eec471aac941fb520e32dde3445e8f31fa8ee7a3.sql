SET allow_experimental_insert_into_iceberg = 1;

INSERT INTO output_iceberg
SELECT
    url,
    count() AS cnt
FROM hits_clickhouse
GROUP BY url
ORDER BY cnt DESC
