CREATE TABLE all_types_test (id UInt32, Date_value Date) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, '2024-01-15');
CREATE DICTIONARY all_types_dict (id UInt32, Date_value Date) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
-- for key which exists
SELECT dictGetDate('all_types_dict', 'Date_value', 1);

-- for key which does not exist, returns the provided default value
SELECT dictGetDateOrDefault('all_types_dict', 'Date_value', 999, toDate('1970-01-01'));
