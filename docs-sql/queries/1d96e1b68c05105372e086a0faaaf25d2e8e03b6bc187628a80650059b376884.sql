CREATE TABLE [IF NOT EXISTS] [db.]table_name [ON CLUSTER cluster]
(
    name1 type1 [DEFAULT|MATERIALIZED|ALIAS expr1],
    name2 type2 [DEFAULT|MATERIALIZED|ALIAS expr2],
    ...
) ENGINE = PostgreSQL({host:port, database, table, user, password[, schema, [, on_conflict]] | named_collection[, option=value [,..]]})
SETTINGS
    [ postgresql_connection_pool_size=16, ]
    [ postgresql_connection_pool_wait_timeout=5000, ]
    [ postgresql_connection_pool_retries=2, ]
    [ postgresql_connection_pool_auto_close_connection=false, ]
    [ postgresql_connection_attempt_timeout=2 ]
;
