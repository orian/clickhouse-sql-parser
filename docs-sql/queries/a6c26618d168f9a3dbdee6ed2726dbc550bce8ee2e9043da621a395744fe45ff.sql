ClickHouse local version 26.3.1.1.

:) ls

SELECT _file AS file
FROM file('*', 'One')
ORDER BY file ASC
