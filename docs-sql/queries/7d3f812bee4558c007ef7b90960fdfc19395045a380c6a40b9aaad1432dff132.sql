SELECT id, name, status, error, start_time, end_time, num_files, uncompressed_size, compressed_size
FROM system.backups
WHERE id = 'events_snapshot_1'
FORMAT Vertical
