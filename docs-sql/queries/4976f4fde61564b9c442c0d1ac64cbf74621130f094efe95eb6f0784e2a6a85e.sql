INSERT INTO ontime SELECT * FROM s3('https://clickhouse-public-datasets.s3.amazonaws.com/ontime/csv_by_year/*.csv.gz', NOSIGN, CSVWithNames) SETTINGS max_insert_threads = 40;
