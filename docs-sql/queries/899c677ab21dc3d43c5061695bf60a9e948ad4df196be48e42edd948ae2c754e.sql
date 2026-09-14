SELECT *
FROM s3(
  'https://<my-bucket>.s3.<my-region>.amazonaws.com/path/to/my-data.parquet.snappy',
  compression_format = 'none'
);
