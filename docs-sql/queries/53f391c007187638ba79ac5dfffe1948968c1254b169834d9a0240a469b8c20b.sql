CREATE TABLE all_types_test (id UInt32, Float64_value Float64) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, 987654.123456);
CREATE DICTIONARY all_types_dict (id UInt32, Float64_value Float64) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
-- for key which exists
SELECT dictGetFloat64('all_types_dict', 'Float64_value', 1);

-- for key which does not exist, returns the provided default value (nan)
SELECT dictGetFloat64OrDefault('all_types_dict', 'Float64_value', 999, nan);
