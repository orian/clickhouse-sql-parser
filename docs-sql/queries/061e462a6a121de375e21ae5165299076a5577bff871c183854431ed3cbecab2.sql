INSERT INTO sometable
FROM INFILE 'data.clickhouse'
COMPRESSION 'lz4'
FORMAT Native
