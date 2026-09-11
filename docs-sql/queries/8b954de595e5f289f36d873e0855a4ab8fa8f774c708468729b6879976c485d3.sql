ALTER TABLE t
MODIFY PROJECTION p (
    SELECT x ORDER BY x
) WITH SETTINGS (
    index_granularity = 128
);
