   CREATE DATABASE demodb
   ENGINE = MaterializedPostgreSQL('postgres1:5432', 'postgres_database', 'postgres_user', 'postgres_password')
   SETTINGS
     materialized_postgresql_replication_slot = 'clickhouse_sync',
     materialized_postgresql_snapshot = '0000000A-0000023F-3',
     materialized_postgresql_tables_list = 'table1,table2,table3';
