BACKUP ALL
EXCEPT TABLES mydb.staging_table
TO S3('https://my-bucket.s3.us-east-1.amazonaws.com/snapshots/full/', 'ACCESS_KEY_ID', 'SECRET_ACCESS_KEY')
SETTINGS experimental_lightweight_snapshot = true, id = 'full_snapshot_1'
