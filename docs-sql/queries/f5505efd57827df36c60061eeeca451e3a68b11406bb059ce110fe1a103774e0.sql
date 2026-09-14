CREATE TABLE all_types_test (id UInt32, UInt16_value UInt16) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, 5000);
CREATE DICTIONARY all_types_dict (id UInt32, UInt16_value UInt16) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
SELECT dictGetUInt16('all_types_dict', 'UInt16_value', 1)
