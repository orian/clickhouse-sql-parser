CREATE TABLE events
    ENGINE = IcebergS3('https://my-bucket.s3.amazonaws.com/warehouse/events/')

SELECT count() FROM events WHERE event_date = today()
