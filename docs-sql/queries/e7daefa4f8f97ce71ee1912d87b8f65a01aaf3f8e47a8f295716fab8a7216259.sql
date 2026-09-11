SELECT name
FROM system.database_engines
WHERE name IN ('Atomic', 'Ordinary')
ORDER BY name
