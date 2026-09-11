CREATE TABLE trips (id UInt64, rider String, fare Float64)
ENGINE = MergeTree
ORDER BY id
SETTINGS storage_policy = 's3_main';
