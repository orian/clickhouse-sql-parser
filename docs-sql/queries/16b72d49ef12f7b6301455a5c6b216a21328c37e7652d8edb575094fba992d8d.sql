CREATE TABLE my_clickhouse_table
(
    id Int64,
    name String
)
ENGINE = MergeTree
ORDER BY id;

INSERT INTO my_clickhouse_table
SELECT * FROM horizon_catalog.`PUBLIC.test_table`;
