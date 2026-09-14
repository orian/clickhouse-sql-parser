SELECT * FROM icebergS3('https://{your-bucket}.s3.{region}.amazonaws.com/{iceberg-path}/', 'NOSIGN')
LIMIT 5;
