CREATE TABLE t (num UInt8) ENGINE = Memory;
INSERT INTO t VALUES (1), (1), (2), (2), (3);

SELECT count() FROM t
