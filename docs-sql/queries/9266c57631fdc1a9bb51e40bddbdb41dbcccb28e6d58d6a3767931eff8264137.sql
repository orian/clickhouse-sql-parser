select * from {{ clickhouse_s3source('taxi_s3', path='/trips_4.gz') }}
