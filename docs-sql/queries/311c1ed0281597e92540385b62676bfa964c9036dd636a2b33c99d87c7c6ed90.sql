SELECT count() FROM table WHERE mapContainsKey(map, 'clickhouse');
-- O
SELECT count() FROM table WHERE mapContains(map, 'clickhouse');
