SELECT *
FROM file('data.parquet', Parquet)
LIMIT 3;
