CREATE TABLE foo (id UInt64) ENGINE = MergeTree ORDER BY id PARTITION BY toYYYYMM(d) SETTINGS index_granularity=8192, parts_to_throw_insert=300
