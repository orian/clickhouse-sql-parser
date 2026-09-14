CREATE TABLE s3queue_engine_table (name String, value UInt32)
ENGINE=S3Queue('https://clickhouse-public-datasets.s3.amazonaws.com/my-test-bucket-768/*', 'CSV', 'gzip')
SETTINGS
    mode = 'unordered',
    after_processing = 'move',
    after_processing_retries = 20,
    after_processing_move_prefix = 'dst_prefix',
    after_processing_move_uri = 'https://clickhouse-public-datasets.s3.amazonaws.com/dst-bucket',
    after_processing_move_access_key_id = 'test',
    after_processing_move_secret_access_key = 'test';
