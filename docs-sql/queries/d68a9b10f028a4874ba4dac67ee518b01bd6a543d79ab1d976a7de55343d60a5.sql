SELECT *
FROM file('dump.sql', MySQLDump)
LIMIT 5
