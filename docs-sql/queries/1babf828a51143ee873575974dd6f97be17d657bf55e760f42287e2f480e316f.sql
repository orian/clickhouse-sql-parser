CREATE TABLE test (d Dynamic) ENGINE = Memory;
INSERT INTO test VALUES (NULL), (42), ('42.42'), (true), ('e10');
SELECT d::Nullable(Float64) FROM test;
