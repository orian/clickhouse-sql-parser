CREATE TABLE test (a Int32, b String)
ENGINE = MergeTree() ORDER BY a
SETTINGS disk = disk(
    type = s3,
    endpoint = 'https://s3.eu-west-1.amazonaws.com/clickhouse-eu-west-1.clickhouse.com/data/',
    use_environment_credentials = 1
);
