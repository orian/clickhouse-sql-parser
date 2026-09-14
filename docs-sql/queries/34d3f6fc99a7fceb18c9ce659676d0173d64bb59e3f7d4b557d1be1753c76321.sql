SELECT count() FROM s3('https://datasets-documentation.s3.eu-west-3.amazonaws.com/mta/*.tsv','CSVWithNames',extra_credentials(role_arn = 'arn:aws:iam::111111111111:role/ClickHouseAccessRole-001'))
