DROP TABLE IF EXISTS test_json;
CREATE TABLE test_json(json JSON(a UInt32)) ENGINE = Memory;
INSERT INTO test_json VALUES ('{"b" : "Hello"}'), ('{"b" : "World", "c" : [1, 2, 3]}');

SELECT distinctJSONPathsAndTypes(json) FROM test_json;
