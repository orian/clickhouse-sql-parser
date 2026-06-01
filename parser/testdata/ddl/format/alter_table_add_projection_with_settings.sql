-- Origin SQL:
ALTER TABLE t ADD PROJECTION p (SELECT x ORDER BY x) WITH SETTINGS (index_granularity = 4096);

ALTER TABLE t ADD PROJECTION IF NOT EXISTS p2 (SELECT x, y ORDER BY x) WITH SETTINGS (index_granularity = 8192, index_granularity_bytes = 1048576) AFTER p;


-- Format SQL:
ALTER TABLE t ADD PROJECTION p (SELECT x ORDER BY x) WITH SETTINGS (index_granularity=4096);
ALTER TABLE t ADD PROJECTION IF NOT EXISTS p2 (SELECT x, y ORDER BY x) WITH SETTINGS (index_granularity=8192, index_granularity_bytes=1048576) AFTER p;
