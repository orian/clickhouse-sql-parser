CREATE TABLE my_s3_table(name String, value UInt32)
ENGINE = S3('https://my-bucket.s3.amazonaws.com/data/*.csv', extra_credentials(role_arn = 'arn:aws:iam::111111111111:role/ClickHouseAccessRole-001'), 'CSV')
