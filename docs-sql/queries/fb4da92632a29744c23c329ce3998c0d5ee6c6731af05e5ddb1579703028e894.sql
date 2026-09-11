CREATE TABLE all_types_test (id UInt32, Float32_value Float32) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, -123.123);
CREATE DICTIONARY all_types_dict (id UInt32, Float32_value Float32) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
SELECT dictGetFloat32('all_types_dict', 'Float32_value', 1)
