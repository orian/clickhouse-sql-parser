INSERT INTO hackernews SELECT * FROM s3('https://clickhouse-datasets.s3.amazonaws.com/hackernews-miniLM/hackernews_part_1_of_1.parquet', NOSIGN);
