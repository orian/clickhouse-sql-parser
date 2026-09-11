CREATE TABLE imported_from_parquet
ENGINE = MergeTree
ORDER BY tuple() AS
SELECT *
FROM file('data.parquet', Parquet)
