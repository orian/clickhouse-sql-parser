SELECT count() FROM table WHERE mapContainsKey(map, 'clickhouse');
-- 또는
SELECT count() FROM table WHERE mapContains(map, 'clickhouse');
