    BACKUP DATABASE test_backups 
    TO S3(
      'https://testchbackups.s3.amazonaws.com/<uuid>',
      '<key id>',
      '<key secret>'
    )
