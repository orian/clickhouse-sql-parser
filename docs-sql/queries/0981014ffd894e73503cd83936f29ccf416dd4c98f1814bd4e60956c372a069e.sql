BACKUP TABLE mydb.events
TO S3('https://my-bucket.s3.us-east-1.amazonaws.com/snapshots/events/', 'ACCESS_KEY_ID', 'SECRET_ACCESS_KEY')
SETTINGS experimental_lightweight_snapshot = true, id = 'events_snapshot_1'
