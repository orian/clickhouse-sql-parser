BACKUP DATABASE my_database 
TO S3('https://your-bucket.s3.amazonaws.com/backup.zip', 'key', 'secret')
ASYNC;
       
-- Returns immediately with backup ID
-- Example result:
-- ┌─id──────────────────────────────────┬─status────────────┐
-- │ abc123-def456-789                   │ CREATING_BACKUP   │
-- └─────────────────────────────────────┴───────────────────┘
