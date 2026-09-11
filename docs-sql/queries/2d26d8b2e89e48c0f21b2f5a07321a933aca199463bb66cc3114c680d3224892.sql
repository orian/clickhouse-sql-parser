RESTORE TABLE test_db.test_table
FROM Disk('backups', 'password-protected.zip')
SETTINGS password='qwerty'
