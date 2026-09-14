CREATE TABLE ext_dict_test_source (id UInt32, c1 UInt32, c2 String) ENGINE = Memory;
INSERT INTO ext_dict_test_source VALUES (1, 1, '1'), (2, 2, '2'), (3, 3, '3');
CREATE DICTIONARY ext_dict_test (id UInt32, c1 UInt32, c2 String) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'ext_dict_test_source' DB currentDatabase())) LAYOUT(FLAT()) LIFETIME(MIN 0 MAX 0);
SELECT dictGet('ext_dict_test', 'c1', toUInt64(1)) AS val
