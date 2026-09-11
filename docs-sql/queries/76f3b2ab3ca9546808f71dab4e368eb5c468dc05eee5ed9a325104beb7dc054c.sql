CREATE TABLE s3_engine_table (name String, value UInt32)
    ENGINE = S3(path, [NOSIGN | aws_access_key_id, aws_secret_access_key,] format, [compression])
    [SETTINGS ...]
