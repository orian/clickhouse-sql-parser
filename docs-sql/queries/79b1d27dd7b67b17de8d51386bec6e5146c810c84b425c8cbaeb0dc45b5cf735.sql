CREATE TABLE all_types_test (id UInt32, Int64_value Int64) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, -9223372036854775808);
CREATE DICTIONARY all_types_dict (id UInt32, Int64_value Int64) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
-- for key which exists
SELECT dictGetInt64('all_types_dict', 'Int64_value', 1);

-- for key which does not exist, returns the provided default value (-1)
SELECT dictGetInt64OrDefault('all_types_dict', 'Int64_value', 999, -1);
