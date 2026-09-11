DESCRIBE TABLE s3('https://datasets-documentation.s3.eu-west-3.amazonaws.com/stackoverflow/parquet/posts/*.parquet', NOSIGN)
SETTINGS describe_compact_output = 1
