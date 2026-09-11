INSERT INTO hits_clickhouse
SELECT url, eventtime, counterid
FROM hits_iceberg
