SELECT name, syntax
FROM system.dictionary_sources
WHERE name IN ('clickhouse', 'file')
ORDER BY name
