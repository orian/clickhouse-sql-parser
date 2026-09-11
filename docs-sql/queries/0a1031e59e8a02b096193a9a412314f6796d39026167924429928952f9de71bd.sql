CREATE TEMPORARY TABLE t_src (id UInt32, val String);
INSERT INTO t_src VALUES (1, 'a'), (2, 'b');

CREATE TEMPORARY VIEW tview AS
SELECT id, upper(val) AS u
FROM t_src
WHERE id <= 2;

SELECT * FROM tview ORDER BY id;
