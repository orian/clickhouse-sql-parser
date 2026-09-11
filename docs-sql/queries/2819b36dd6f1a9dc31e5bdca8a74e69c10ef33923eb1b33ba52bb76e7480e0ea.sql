SELECT pg_xlog_location_diff(pg_current_xlog_insert_location(), '0/0') / 1024 / 1024 AS wal_generated_mb;
