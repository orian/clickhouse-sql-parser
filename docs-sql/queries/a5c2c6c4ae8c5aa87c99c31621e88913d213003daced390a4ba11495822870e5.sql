SELECT
    name as mv_name,
    replaceRegexpOne(
        create_table_query,
        '.*TO\\s+`?([^`\\s(]+)`?\\.`?([^`\\s(]+)`?.*',
        '\\1.\\2'
    ) AS target_table
FROM system.tables
WHERE database = 'your_schema'
  AND engine = 'MaterializedView'
