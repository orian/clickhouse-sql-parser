CREATE TABLE random (a Array(Int8), d Decimal32(4), c Tuple(DateTime64(3), UUID)) ENGINE=Memory;
INSERT INTO random SELECT * FROM generateRandom() LIMIT 2;
SELECT * FROM random;
