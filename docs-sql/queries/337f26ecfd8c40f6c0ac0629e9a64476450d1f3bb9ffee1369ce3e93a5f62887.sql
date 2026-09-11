CREATE TABLE t (UserID UInt32) ENGINE = Memory;
INSERT INTO t VALUES (1), (1), (2), (3);

SELECT groupBitmap(UserID) AS num FROM t;
