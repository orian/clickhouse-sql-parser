SELECT * FROM s3('https://clickhouse-public-datasets.s3.amazonaws.com/my-test-bucket-768/**', NOSIGN, 'CSV', 'name String, value UInt32', 'gzip');
