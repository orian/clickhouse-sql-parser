ALTER TABLE t
ADD PROJECTION p (
    SELECT x ORDER BY x
) WITH SETTINGS (
    index_granularity = 4096,
    index_granularity_bytes = 1048576
);
