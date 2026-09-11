CREATE TABLE all_types_test (id UInt32, UUID_value UUID) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, '550e8400-e29b-41d4-a716-446655440000');
CREATE DICTIONARY all_types_dict (id UInt32, UUID_value UUID) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
-- for key which exists
SELECT dictGetUUID('all_types_dict', 'UUID_value', 1);

-- for key which does not exist, returns the provided default value
SELECT dictGetUUIDOrDefault('all_types_dict', 'UUID_value', 999, '00000000-0000-0000-0000-000000000000'::UUID);
