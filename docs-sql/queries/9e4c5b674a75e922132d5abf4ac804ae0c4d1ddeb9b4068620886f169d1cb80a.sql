    SELECT throwIf((
       SELECT groupArray(tuple(*))
       FROM test_db.test_table
       ) != (
       SELECT groupArray(tuple(*))
       FROM test_db.test_table_restored
    ), 'Data does not match after BACKUP/RESTORE')
