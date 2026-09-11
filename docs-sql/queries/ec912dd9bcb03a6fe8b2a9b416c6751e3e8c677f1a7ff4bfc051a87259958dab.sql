RESTORE DATABASE nyc_taxi
FROM S3(
    'BUCKET_URL/backup_s2.zip',
    extra_credentials(role_arn = 'ROLE_ARN')
)
SETTINGS allow_non_empty_tables=true;
