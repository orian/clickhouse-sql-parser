CREATE TABLE s3_engine_table (name String, value UInt32)
    ENGINE = S3(path [, NOSIGN | aws_access_key_id, aws_secret_access_key,] format, [compression], [partition_strategy], [partition_columns_in_data_file], [extra_credentials])
    [PARTITION BY expr]
    [SETTINGS ...]
