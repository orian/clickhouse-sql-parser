CREATE TABLE logs (id bigint, payload bytea) WITH (
    copy_from = 's3://my-bucket/logs.parquet'
);
