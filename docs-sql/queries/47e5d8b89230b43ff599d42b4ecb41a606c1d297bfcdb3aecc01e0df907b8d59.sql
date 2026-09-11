CREATE TABLE myNewTable AS myOldTable;

INSERT INTO myNewTable
SELECT * FROM myOldTable
WHERE CounterID <12345;

DROP TABLE myOldTable;

RENAME TABLE myNewTable TO myOldTable;
