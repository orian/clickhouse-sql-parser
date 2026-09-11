BACKUP DATABASE nyc_taxi
TO S3(
  'BUCKET_URL/backup_s2.zip',
  'KEY_ID',
  'SECRET_KEY'
)
