CREATE TABLE test (d Dynamic) ENGINE=Memory;
INSERT INTO test VALUES (NULL), (1::Int8), (2::Int16), (3::Int32), (4::Int64);
