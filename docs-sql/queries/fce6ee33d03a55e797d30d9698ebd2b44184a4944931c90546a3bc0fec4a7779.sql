SELECT * FROM postgresql('localhost:5432', 'test', 'test', 'postgresql_user', 'password', SETTINGS postgresql_connection_pool_size = 32);
