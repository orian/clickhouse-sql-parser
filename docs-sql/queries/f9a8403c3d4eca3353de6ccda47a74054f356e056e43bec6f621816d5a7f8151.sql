    RESTORE TABLE postgres_db.table1 AS restored_db.table1
    FROM Disk('backups', 'postgres_db.zip')
    SETTINGS allow_different_table_def = 1;
