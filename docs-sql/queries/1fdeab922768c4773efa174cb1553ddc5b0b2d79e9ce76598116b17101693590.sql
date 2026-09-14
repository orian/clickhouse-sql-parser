CREATE TABLE t_null
(
  x Int32,
  y Nullable(Int32)
)
ENGINE = MergeTree
ORDER BY tuple();

INSERT INTO t_null VALUES (1, NULL), (2, 0), (3, 3);

SELECT x FROM t_null WHERE isZeroOrNull(y);
