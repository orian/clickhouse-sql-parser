CREATE TABLE big_table (name String, value UInt32)
    ENGINE = S3('https://clickhouse-public-datasets.s3.amazonaws.com/my-bucket/my_folder/file-{000..999}.csv', 'CSV');
