SOURCE(CASSANDRA(
    host 'localhost'
    port 9042
    user 'username'
    password 'qwerty123'
    keyspace 'database_name'
    column_family 'table_name'
    allow_filtering 1
    partition_key_prefix 1
    consistency 'One'
    where '"SomeColumn" = 42'
    max_threads 8
    query 'SELECT id, value_1, value_2 FROM database_name.table_name'
))
