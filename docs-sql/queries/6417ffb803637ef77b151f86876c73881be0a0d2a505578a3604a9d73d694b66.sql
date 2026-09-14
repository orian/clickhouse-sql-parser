SELECT comment
FROM system.tables
WHERE database = currentDatabase() AND name = 'table_with_comment';
