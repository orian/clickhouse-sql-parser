CREATE TABLE table_from_mysql
ENGINE = MergeTree
ORDER BY tuple() AS
SELECT *
FROM file('mysql.sql', MySQLDump)
