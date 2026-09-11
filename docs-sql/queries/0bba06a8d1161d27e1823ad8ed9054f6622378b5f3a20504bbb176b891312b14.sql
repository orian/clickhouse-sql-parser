CREATE TABLE all_types_test (id UInt32, UInt32_value UInt32) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, 1000000);
CREATE DICTIONARY all_types_dict (id UInt32, UInt32_value UInt32) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
SELECT dictGetUInt32('all_types_dict', 'UInt32_value', 1)
