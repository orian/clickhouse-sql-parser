CREATE TABLE all_types_test (id UInt32, String_value String) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, 'test string');
CREATE DICTIONARY all_types_dict (id UInt32, String_value String) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
SELECT dictGetString('all_types_dict', 'String_value', 1)
