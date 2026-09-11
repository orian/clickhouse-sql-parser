INSERT INTO FUNCTION s3(
    'http://seaweedfs:8333/clickhouse/sample/trips.parquet',
    'your_access_key_id',
    'your_secret_access_key',
    'Parquet'
)
SELECT number AS id, concat('rider_', toString(number % 10)) AS rider, number * 1.5 AS fare
FROM numbers(1000);

SELECT count()
FROM s3(
    'http://seaweedfs:8333/clickhouse/sample/*.parquet',
    'your_access_key_id',
    'your_secret_access_key',
    'Parquet'
);
