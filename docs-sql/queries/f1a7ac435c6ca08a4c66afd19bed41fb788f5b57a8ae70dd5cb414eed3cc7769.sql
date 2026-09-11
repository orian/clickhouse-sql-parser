SELECT arrayJoin(distinctJSONPathsAndTypes(json))
FROM s3('s3://clickhouse-public-datasets/gharchive/original/2020-01-01-*.json.gz', NOSIGN, JSONAsObject)
SETTINGS date_time_input_format = 'best_effort'
