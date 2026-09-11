CREATE TABLE all_types_test (id UInt32, Int8_value Int8) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, -100);
CREATE DICTIONARY all_types_dict (id UInt32, Int8_value Int8) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
SELECT dictGetInt8('all_types_dict', 'Int8_value', 1)
