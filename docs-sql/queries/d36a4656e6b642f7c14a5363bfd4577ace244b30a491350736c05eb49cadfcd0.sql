SELECT *
FROM s3(
    'https://your-bucket.t3.storage.dev/path/to/file.parquet',
    'your_access_key_id',
    'your_secret_access_key',
    'Parquet'
);
