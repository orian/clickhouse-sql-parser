DESCRIBE TABLE s3('https://datasets-documentation.s3.eu-west-3.amazonaws.com/json/conflict_sample.json', NOSIGN, JSONAsObject)
SETTINGS enable_json_type = 1, describe_compact_output = 1
