SET join_runtime_filter_min_probe_rows = 10;

CREATE TABLE t1 (id UInt64, value String) ENGINE = MergeTree ORDER BY id;
CREATE TABLE t2 (id UInt64, value String) ENGINE = MergeTree ORDER BY id;
INSERT INTO t1 SELECT number, toString(number) FROM numbers(100);
INSERT INTO t2 SELECT number, toString(number) FROM numbers(100);

EXPLAIN actions = 1, compact = 1, pretty = 1
SELECT * FROM t1 INNER JOIN t2 ON t1.id = t2.id FORMAT Raw;
