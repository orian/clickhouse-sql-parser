    CREATE TABLE output_iceberg
    (
        url String,
        cnt UInt64
    )
    ENGINE = IcebergS3('https://your-bucket.s3.amazonaws.com/output/', 'access_key', 'secret_key')
