CREATE TABLE secret_data (id UInt64) ENGINE = MergeTree ORDER BY id
SETTINGS storage_policy = 'encrypted';
