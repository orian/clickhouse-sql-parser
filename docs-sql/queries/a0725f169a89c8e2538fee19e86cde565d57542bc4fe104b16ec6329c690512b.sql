CREATE DATABASE mysql_db
ENGINE = MySQL('localhost:3306', 'test', 'my_user', 'user_password')
SETTINGS enable_compression = 1;
