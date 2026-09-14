CREATE TABLE mysql_engine_compression
(
    id UInt32,
    name String,
    age UInt32,
    money UInt32
)
ENGINE = MySQL('mysql80:3306', 'clickhouse', 'test_table', 'root', 'password')
SETTINGS enable_compression = 1;
