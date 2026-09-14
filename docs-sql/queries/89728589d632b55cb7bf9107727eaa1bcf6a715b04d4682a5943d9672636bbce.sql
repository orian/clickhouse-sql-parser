INSERT INTO FUNCTION s3(s3_conn, filename='t_03363_function', format=Parquet, partition_strategy='hive') PARTITION BY (year, country) SELECT 2020 as year, 'Russia' as country, 1 as id;
