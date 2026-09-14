CREATE TABLE table_with_asterisk (name String, value UInt32)
    ENGINE = S3('https://clickhouse-public-datasets.s3.amazonaws.com/my-bucket/{some,another}_folder/*', 'CSV');
