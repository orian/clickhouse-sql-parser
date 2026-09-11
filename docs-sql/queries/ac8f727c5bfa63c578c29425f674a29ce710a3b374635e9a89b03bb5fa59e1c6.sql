CREATE TABLE example (values Tuple(name String, age UInt32)) ENGINE = Memory;
INSERT INTO example VALUES (('Alice', 30));
SELECT tupleElement(values, 'name') FROM example;
