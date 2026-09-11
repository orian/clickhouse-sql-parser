CREATE TABLE table_with_question_mark (name String, value UInt32)
    ENGINE = S3('https://clickhouse-public-datasets.s3.amazonaws.com/my-bucket/{some,another}_folder/some_file_?', 'CSV');
