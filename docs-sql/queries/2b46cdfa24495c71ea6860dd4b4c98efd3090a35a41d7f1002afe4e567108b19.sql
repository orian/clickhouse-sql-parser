RESTORE TABLE test_db.test_table AS test_db.test_table2
FROM Disk('backups', 'incremental-a.zip');
