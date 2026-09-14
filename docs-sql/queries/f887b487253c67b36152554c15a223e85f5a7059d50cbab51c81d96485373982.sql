CREATE TABLE all_types_test (id UInt32, UInt64_value UInt64) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, 9223372036854775807);
CREATE DICTIONARY all_types_dict (id UInt32, UInt64_value UInt64) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
-- for key which exists
SELECT dictGetUInt64('all_types_dict', 'UInt64_value', 1);

-- for key which does not exist, returns the provideddefault value (0)
SELECT dictGetUInt64OrDefault('all_types_dict', 'UInt64_value', 999, 0);
