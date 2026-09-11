SELECT *
FROM file(dump.sql, MySQLDump)
SETTINGS input_format_mysql_dump_table_name = 'test2'
