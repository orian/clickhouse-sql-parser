    RESTORE TABLE data AS test_db.test_table_restored FROM S3(
    'https://backup-ch-docs.s3.us-east-1.amazonaws.com/backups/incremental_backup',
    '<access key id>',
    '<secret access key>'
    )
