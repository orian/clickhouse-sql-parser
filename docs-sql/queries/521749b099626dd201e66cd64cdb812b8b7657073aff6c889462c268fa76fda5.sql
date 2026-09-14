CREATE TABLE all_types_test (id UInt32, Int64_value Int64) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, -9223372036854775807);
CREATE DICTIONARY all_types_dict (id UInt32, Int64_value Int64) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
SELECT dictGetInt64('all_types_dict', 'Int64_value', 1)
