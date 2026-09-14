SELECT count() FROM table WHERE mapContainsKey(map, 'clickhouse');
-- OU
SELECT count() FROM table WHERE mapContains(map, 'clickhouse');
