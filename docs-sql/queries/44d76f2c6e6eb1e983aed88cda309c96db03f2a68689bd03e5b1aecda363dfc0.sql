RESTORE
    TABLE system.users,
    TABLE system.roles,
    TABLE system.settings_profiles,
    TABLE system.row_policies,
    TABLE system.quotas,
    ALL EXCEPT DATABASES INFORMATION_SCHEMA, information_schema, system
FROM S3(
    'BUCKET_URL',
    extra_credentials(role_arn = 'ROLE_ARN')
)
