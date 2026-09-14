CREATE TABLE test (d Dynamic) ENGINE = Memory;
INSERT INTO test VALUES (NULL), (42), ('Hello, World!'), ([1, 2, 3]);
SELECT d, dynamicType(d), dynamicElement(d, 'String'), dynamicElement(d, 'Int64'), dynamicElement(d, 'Array(Int64)'), dynamicElement(d, 'Date'), dynamicElement(d, 'Array(String)') FROM test
