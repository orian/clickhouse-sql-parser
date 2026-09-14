CREATE TABLE events
    ENGINE = IcebergS3('https://my-bucket.s3.amazonaws.com/warehouse/events/')
SETTINGS iceberg_metadata_async_prefetch_period_ms = 60000;

SELECT count()
FROM events
SETTINGS iceberg_metadata_staleness_ms = 60000;
