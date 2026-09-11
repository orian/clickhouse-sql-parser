CREATE TABLE test (json JSON) ENGINE = Memory;
INSERT INTO test VALUES ('{"a" : 42, "b" : {"c" : 1, "d" : "Hello"}}'), ('{"a" : {"x": 1, "y": 2}, "b" : {"c" : 1}}'), ('{"c" : "World"}');
SELECT json FROM test;
