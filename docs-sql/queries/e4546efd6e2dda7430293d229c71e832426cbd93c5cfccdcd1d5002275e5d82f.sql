SELECT * FROM chdb_query(
    'SELECT number AS n, number * number FROM numbers(5) ORDER BY n'
) AS (n int2, p int);
