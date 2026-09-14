CREATE TABLE p
(
    `column1` UInt32,
    `column2` UInt32,
    `column3` UInt32
)
ENGINE = S3(
           'http://minio:10000/clickhouse//test_{_partition_id}.csv',
           'minioadmin',
           'minioadminpassword',
           'CSV',
           partition_strategy='wildcard')
PARTITION BY column3
