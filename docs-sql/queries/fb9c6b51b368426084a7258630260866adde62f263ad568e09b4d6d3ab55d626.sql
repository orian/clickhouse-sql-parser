CREATE TABLE t_null (x Int8, y Nullable(Int8))
ENGINE=MergeTree()
ORDER BY x;

INSERT INTO t_null VALUES (1, NULL), (2, 3);

SELECT assumeNotNull(y) FROM t_null;
SELECT toTypeName(assumeNotNull(y)) FROM t_null;
