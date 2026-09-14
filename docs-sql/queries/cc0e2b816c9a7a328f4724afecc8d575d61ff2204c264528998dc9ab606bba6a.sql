-- Remove orphan files older than a specific timestamp
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('2026-03-01 00:00:00');

-- Dry run: preview which files would be deleted
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(dry_run = 1);

-- Scan only the data directory
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = '2026-03-01 00:00:00',
    location = 'data/'
);

-- Combine positional older_than with named arguments
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    '2026-03-01 00:00:00',
    dry_run = 1
);
