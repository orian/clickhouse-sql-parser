SELECT *
FROM s3('http://minio:10000/clickhouse//test_3.csv', 'minioadmin', 'minioadminpassword', 'CSV')
