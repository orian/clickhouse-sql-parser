SYSTEM RESTORE REPLICA table_name;  -- It will detach all partitions, re-create metadata in ZK (like it's new empty table), and then attach all partitions back

SYSTEM SYNC REPLICA table_name; -- Wait for replicas to synchronize parts. Also it's recommended to check `system.detached_parts` on all replicas after recovery is finished.
