CREATE TABLE all_types_test (id UInt32, DateTime_value DateTime) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, '2024-01-15 10:30:00');
CREATE DICTIONARY all_types_dict (id UInt32, DateTime_value DateTime) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
SELECT dictGetDateTime('all_types_dict', 'DateTime_value', 1)
