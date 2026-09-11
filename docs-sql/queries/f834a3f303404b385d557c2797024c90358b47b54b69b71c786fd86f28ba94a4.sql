CREATE SERVER taxi_srv FOREIGN DATA WRAPPER clickhouse_fdw
       OPTIONS(driver 'binary', host 'localhost', dbname 'taxi');
