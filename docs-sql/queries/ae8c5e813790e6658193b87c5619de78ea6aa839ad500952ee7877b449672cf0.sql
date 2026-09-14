SELECT name, syntax
FROM system.data_skipping_index_types
WHERE name IN ('minmax', 'set')
ORDER BY name
