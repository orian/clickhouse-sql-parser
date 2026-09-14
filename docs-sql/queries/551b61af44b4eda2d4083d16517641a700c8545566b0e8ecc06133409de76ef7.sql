SELECT count()
FROM icebergS3('https://my-bucket.s3.amazonaws.com/warehouse/my_table/')
WHERE event_date >= today() - 7
