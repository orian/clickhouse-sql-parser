CREATE TABLE all_types_test (id UInt32, IPv6_value IPv6) ENGINE = MergeTree() ORDER BY id;
INSERT INTO all_types_test VALUES (1, '2001:db8:85a3::8a2e:370:7334');
CREATE DICTIONARY all_types_dict (id UInt32, IPv6_value IPv6) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'all_types_test' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 300 MAX 600);
-- for key which exists
SELECT dictGetIPv6('all_types_dict', 'IPv6_value', 1);

-- for key which does not exist, returns the provided default value
SELECT dictGetIPv6OrDefault('all_types_dict', 'IPv6_value', 999, '::1'::IPv6);
