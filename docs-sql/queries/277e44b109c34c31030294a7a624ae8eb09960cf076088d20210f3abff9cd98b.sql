CREATE TABLE some_table (id1 UInt32, id2 UInt32, name String) ENGINE = Join(ANY, LEFT, id1, id2);
INSERT INTO some_table VALUES (1, 11, 'a') (2, 12, 'b') (3, 13, 'c');

SELECT joinGet(some_table, 'name', toUInt32(1), toUInt32(11));
