    CREATE DATABASE IF NOT EXISTS test_db;
    CREATE TABLE test_db.test_table
    (
        `key` Int,
        `value` String,
        `array` Array(String)
    )
    ENGINE = MergeTree
    ORDER BY tuple()
