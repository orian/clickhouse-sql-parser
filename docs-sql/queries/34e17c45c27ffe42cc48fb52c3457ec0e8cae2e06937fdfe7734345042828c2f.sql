SELECT count() FROM table WHERE mapContainsKey(map, 'clickhouse');
-- 或
SELECT count() FROM table WHERE mapContains(map, 'clickhouse');
