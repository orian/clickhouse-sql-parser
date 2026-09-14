CREATE TABLE range_key_dictionary_source_table (key UInt64, start_date Date, end_date Date, value String, value_nullable Nullable(String)) ENGINE = TinyLog();
INSERT INTO range_key_dictionary_source_table VALUES(1, toDate('2019-05-20'), toDate('2019-05-20'), 'First', 'First');
INSERT INTO range_key_dictionary_source_table VALUES(2, toDate('2019-05-20'), toDate('2019-05-20'), 'Second', NULL);
INSERT INTO range_key_dictionary_source_table VALUES(3, toDate('2019-05-20'), toDate('2019-05-20'), 'Third', 'Third');
CREATE DICTIONARY range_key_dictionary (key UInt64, start_date Date, end_date Date, value String, value_nullable Nullable(String)) PRIMARY KEY key SOURCE(CLICKHOUSE(TABLE 'range_key_dictionary_source_table' DB currentDatabase())) LIFETIME(MIN 1 MAX 1000) LAYOUT(RANGE_HASHED()) RANGE(MIN start_date MAX end_date);
SELECT
    (number, toDate('2019-05-20')),
    dictGetOrNull('range_key_dictionary', 'value', number, toDate('2019-05-20')),
FROM system.numbers LIMIT 5 FORMAT TabSeparated;
