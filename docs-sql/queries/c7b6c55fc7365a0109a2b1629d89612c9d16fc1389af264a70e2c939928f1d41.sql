SELECT count(*)
FROM icebergS3(
  'https://{your-bucket}.s3.{region}.amazonaws.com/{iceberg-path}/',
  'NOSIGN',
  extra_credentials(role_arn='arn:aws:iam::{account-id}:role/iceberg-role-for-clickhouse', role_session_name='iceberg-session')
);
