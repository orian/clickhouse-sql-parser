CREATE TABLE t (s String) ENGINE = Memory;
INSERT INTO t VALUES ('a1'), ('a02');

SELECT s FROM t ORDER BY naturalSortKey(s);
