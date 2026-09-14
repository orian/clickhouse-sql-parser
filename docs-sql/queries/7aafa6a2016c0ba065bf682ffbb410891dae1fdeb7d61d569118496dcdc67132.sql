CREATE TABLE dict_mult_source (id UInt32, c1 UInt32, c2 String) ENGINE = Memory;
INSERT INTO dict_mult_source VALUES (1, 1, '1'), (2, 2, '2'), (3, 3, '3');
CREATE DICTIONARY ext_dict_mult (id UInt32, c1 UInt32, c2 String) PRIMARY KEY id SOURCE(CLICKHOUSE(TABLE 'dict_mult_source' DB currentDatabase())) LAYOUT(FLAT()) LIFETIME(MIN 0 MAX 0);
SELECT
    dictGet('ext_dict_mult', ('c1','c2'), number + 1) AS val,
    toTypeName(val) AS type
FROM system.numbers
LIMIT 3;
