CREATE TABLE pg_table
(
    `float_nullable` Nullable(Float32),
    `str` String,
    `int_id` Int32
)
ENGINE = PostgreSQL('localhost:5432', 'public', 'test', 'postgres_user', 'postgres_password')
SETTINGS postgresql_connection_pool_size = 32, postgresql_connection_pool_auto_close_connection = 1;
