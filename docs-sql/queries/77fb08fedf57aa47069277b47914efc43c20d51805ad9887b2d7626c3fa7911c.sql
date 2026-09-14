CREATE TABLE all_types_test (id UInt32, UInt8_value UInt8) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, 100);
CREATE DICTIONARY all_types_dict (id UInt32, UInt8_value UInt8) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
-- for key which exists
SELECT dictGetUInt8('all_types_dict', 'UInt8_value', 1);

-- for key which does not exist, returns the provided default value (0)
SELECT dictGetUInt8OrDefault('all_types_dict', 'UInt8_value', 999, 0);
