SET allow_experimental_paimon_storage_engine = 1;

-- Local storage
CREATE TABLE paimon_mv_source
ENGINE = PaimonLocal('/path/to/paimon/table')
SETTINGS
    paimon_incremental_read = 1,
    paimon_keeper_path = '/clickhouse/tables/{uuid}',
    paimon_replica_name = '{replica}',
    paimon_metadata_refresh_interval_sec = 1;

-- S3 storage (the `Paimon` engine defaults to the S3 implementation when no `disk` is specified)
CREATE TABLE paimon_mv_source
ENGINE = Paimon('http://minio:9000/bucket/path/to/table', 'access_key', 'secret_key')
SETTINGS
    paimon_incremental_read = 1,
    paimon_keeper_path = '/clickhouse/tables/{uuid}',
    paimon_replica_name = '{replica}',
    paimon_metadata_refresh_interval_sec = 1;
