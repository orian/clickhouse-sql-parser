CREATE TABLE all_types_test (id UInt32, UUID_value UUID) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, '123e4567-e89b-12d3-a456-426614174000');
CREATE DICTIONARY all_types_dict (id UInt32, UUID_value UUID) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
SELECT dictGetUUID('all_types_dict', 'UUID_value', 1)
