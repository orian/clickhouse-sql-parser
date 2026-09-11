SELECT count()
FROM s3('https://clickhouse-public-datasets.s3.amazonaws.com/bluesky/file_0001.json.gz', NOSIGN, 'JSONEachRow')
