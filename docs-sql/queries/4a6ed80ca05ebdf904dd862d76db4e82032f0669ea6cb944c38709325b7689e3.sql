CREATE TABLE all_types_test (id UInt32, Date_value Date) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, '2020-01-01');
CREATE DICTIONARY all_types_dict (id UInt32, Date_value Date) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
SELECT dictGetDate('all_types_dict', 'Date_value', 1)
