RESTORE DATABASE nyc_taxi
FROM S3(
    'BUCKET_URL/backup_s1.zip',
    extra_credentials(role_arn = 'ROLE_ARN')
)
