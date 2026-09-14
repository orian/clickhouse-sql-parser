SELECT *
FROM s3('http://minio:10000/clickhouse//test_45.csv', 'minioadmin', 'minioadminpassword', 'CSV')
