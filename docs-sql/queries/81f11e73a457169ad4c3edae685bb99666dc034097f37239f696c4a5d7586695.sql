CREATE TABLE table_with_range (name String, value UInt32)
    ENGINE = S3('https://clickhouse-public-datasets.s3.amazonaws.com/my-bucket/{some,another}_folder/some_file_{1..3}', 'CSV');
