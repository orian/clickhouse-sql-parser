RESTORE TABLE mydb.events AS mydb.events_restored
FROM S3('https://my-bucket.s3.us-east-1.amazonaws.com/snapshots/events/', 'ACCESS_KEY_ID', 'SECRET_ACCESS_KEY')
SETTINGS snapshot_from_current_service = 1
