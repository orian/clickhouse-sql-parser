INSERT INTO mysql_copy
SELECT * FROM mysql('host:port', 'database', 'table', 'user', 'password')
WHERE id > (SELECT max(id) FROM mysql_copy);
