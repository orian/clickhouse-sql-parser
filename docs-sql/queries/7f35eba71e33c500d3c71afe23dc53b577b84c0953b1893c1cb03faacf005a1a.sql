SELECT * FROM s3('s3://clickhouse-public-datasets/my-test-bucket-768/**/test-data.csv.gz', NOSIGN, 'CSV', 'name String, value UInt32', 'gzip');
