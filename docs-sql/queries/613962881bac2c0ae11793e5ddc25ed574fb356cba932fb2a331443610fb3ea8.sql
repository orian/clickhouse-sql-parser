SELECT * FROM clickhouse_query(
    'server',
    'SELECT id, name, salary FROM remote_table WHERE salary > 50000'
) AS ch(id int, name text, salary numeric);
