SELECT id, info, lock_path
FROM system.snapshot_locks
WHERE id = 'events_snapshot_1'
