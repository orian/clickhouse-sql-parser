INSERT INTO sometable
FROM INFILE 'data.parquet' FORMAT Parquet;

SELECT *
FROM sometable
LIMIT 5;
