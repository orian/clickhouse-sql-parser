    BACKUP TABLE test_db.test_table TO S3(
    'https://backup-ch-docs.s3.us-east-1.amazonaws.com/backups/base_backup',
    '<access key id>',
    '<secret access key>'
    )
