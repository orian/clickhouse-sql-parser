  RESTORE TABLE src AS existing_replacing_mergetree
  FROM Disk('backups', 'table.zip')
  SETTINGS allow_different_table_def = 1;
