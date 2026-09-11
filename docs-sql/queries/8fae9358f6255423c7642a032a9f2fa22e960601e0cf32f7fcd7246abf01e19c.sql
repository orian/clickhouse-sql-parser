SELECT
    n,
    toDateTime(time)                 <--- int to time
FROM file('time.parquet', Parquet);
