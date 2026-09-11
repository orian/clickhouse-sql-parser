-- Positional form: single unnamed older_than argument
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('timestamp')

-- Named form
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = 'timestamp',
    location = 'path',
    dry_run = 0|1
)

-- No arguments: use all defaults (older_than = 3 days ago)
ALTER TABLE iceberg_table EXECUTE remove_orphan_files()
