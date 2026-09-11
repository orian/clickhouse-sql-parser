CREATE TABLE iceberg_secure (
  id UInt64,
  event_date Date,
  data String
)
ENGINE = IcebergS3(
  'https://{your-bucket}.s3.{region}.amazonaws.com/{iceberg-path}/',
  'NOSIGN',
  extra_credentials(role_arn='arn:aws:iam::{account-id}:role/iceberg-role-for-clickhouse')
);
