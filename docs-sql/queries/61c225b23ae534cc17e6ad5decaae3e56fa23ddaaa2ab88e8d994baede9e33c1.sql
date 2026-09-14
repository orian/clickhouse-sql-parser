DESCRIBE TABLE s3('https://datasets-documentation.s3.eu-west-3.amazonaws.com/pypi/pypi_with_tags/sample_rows.json.gz', NOSIGN, 'JSONAsObject')
SETTINGS describe_compact_output = 1
