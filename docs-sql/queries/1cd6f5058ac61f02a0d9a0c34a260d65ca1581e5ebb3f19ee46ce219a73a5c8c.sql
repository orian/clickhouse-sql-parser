DESCRIBE TABLE s3('https://datasets-documentation.s3.eu-west-3.amazonaws.com/pypi/2024-12-17/*.parquet', NOSIGN)
FORMAT PrettyCompactNoEscapesMonoBlock
SETTINGS describe_compact_output = 1
