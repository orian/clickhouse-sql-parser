    CREATE TABLE mysql_table1 (
      id UInt64,
      column1 String
    )
    ENGINE = MySQL('mysql-host.domain.com','db1','table1','mysql_clickhouse','Password123!')
