CREATE TABLE my_db.my_table
ENGINE = IcebergS3(
  's3://{your-bucekt}/warehouse/{db}/{table}/',
  'NOSIGN',
  extra_credentials(role_arn='arn:aws:iam::{account-id}:role/iceberg-role-for-clickhouse')
)
SETTINGS
  catalog_type = 'glue',
  warehouse = '{your-warehouse}',
  storage_endpoint = 's3://{your-bucket}',
  region = '{region}'
  aws_role_arn = 'arn:aws:iam::{account-id}:role/iceberg-role-for-clickhouse';
