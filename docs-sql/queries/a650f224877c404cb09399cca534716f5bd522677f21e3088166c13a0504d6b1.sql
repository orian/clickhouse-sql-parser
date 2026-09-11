SELECT name, description
FROM system.disk_types
WHERE name IN ('local', 'object_storage')
ORDER BY name
