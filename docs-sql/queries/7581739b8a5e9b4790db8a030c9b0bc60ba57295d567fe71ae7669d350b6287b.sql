    CREATE SERVER ch FOREIGN DATA WRAPPER clickhouse_fdw
           OPTIONS(driver 'http', host '<clickhouse_cloud_host>', dbname '<database_name>', port '8443');
