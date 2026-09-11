SELECT *
FROM s3('http://minio:10000/clickhouse//test_3.csv', 'minioadmin', 'minioadminpassword', 'CSV')
SETTINGS filesystem_cache_name = 'cache_for_s3', enable_filesystem_cache = 1;
