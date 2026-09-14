
        CREATE TABLE iceberg_manifest_entries (c0 Int64)
        ENGINE = IcebergLocal((SELECT value FROM system.server_settings WHERE name = 'user_files_path') || '/iceberg_manifest_entries_' || toString(toUnixTimestamp64Micro(now64(6))) || '/')
    