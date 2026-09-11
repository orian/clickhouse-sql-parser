CREATE TABLE ck (id UInt8, name String) ENGINE = Memory;
INSERT INTO ck VALUES (1, 'zhangsan'), (1, 'lisi'), (2, 'wangwu');

SELECT id, groupArray(10)(name) FROM ck GROUP BY id ORDER BY id;
