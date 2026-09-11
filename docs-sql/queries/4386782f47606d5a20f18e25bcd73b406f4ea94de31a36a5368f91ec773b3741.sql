SET input_format_json_read_objects_as_strings = 1;
CREATE TABLE test (id UInt64, obj String, date Date) ENGINE=Memory();
INSERT INTO test FORMAT JSONEachRow {"id" : 1, "obj" : {"a" : 1, "b" : "Hello"}, "date" : "2020-01-01"};
SELECT * FROM test;
