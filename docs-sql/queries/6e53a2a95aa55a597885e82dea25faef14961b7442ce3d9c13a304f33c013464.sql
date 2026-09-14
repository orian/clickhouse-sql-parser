CREATE TABLE all_types_test (id UInt32, IPv4_value IPv4) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, '192.168.0.1');
CREATE DICTIONARY all_types_dict (id UInt32, IPv4_value IPv4) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
SELECT dictGetIPv4('all_types_dict', 'IPv4_value', 1)
