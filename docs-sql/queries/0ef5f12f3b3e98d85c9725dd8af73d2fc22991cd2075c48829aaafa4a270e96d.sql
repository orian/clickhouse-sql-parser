BACKUP { TABLE [db.]table_name | DATABASE db_name | ALL [EXCEPT {TABLES | DATABASES} ...] }
TO { S3(...) | AzureBlobStorage(...) }
SETTINGS experimental_lightweight_snapshot = true, id = '<snapshot_id>'
