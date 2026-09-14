SELECT value, getServerSetting('named_collections_storage_type')
FROM system.server_settings
WHERE name = 'named_collections_storage.type';
