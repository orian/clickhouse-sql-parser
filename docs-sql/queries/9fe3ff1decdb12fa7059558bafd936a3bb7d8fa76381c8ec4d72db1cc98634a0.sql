SELECT *
FROM football
INTO OUTFILE 'football.parquet'
FORMAT Parquet
