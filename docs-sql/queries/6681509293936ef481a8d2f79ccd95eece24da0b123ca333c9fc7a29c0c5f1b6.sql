SET s3_base = 's3://clickhouse-public-datasets/';
SELECT count() FROM s3('hits_compatible/hits.csv', NOSIGN);
