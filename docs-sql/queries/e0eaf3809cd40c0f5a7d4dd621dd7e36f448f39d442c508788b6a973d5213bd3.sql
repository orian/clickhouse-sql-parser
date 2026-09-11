CREATE TABLE prices (item String, total UInt32) ENGINE = Memory;

INSERT INTO prices FORMAT Values ('apple', 3 * 4), ('pear', length('hello') + 10);

SELECT * FROM prices ORDER BY total;
