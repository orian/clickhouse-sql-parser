RESTORE DATABASE test_backups
FROM S3(
  'https://storage.googleapis.com/<bucket>/<uuid>',
  '<hmac-key>',
  '<hmac-secret>'
)
