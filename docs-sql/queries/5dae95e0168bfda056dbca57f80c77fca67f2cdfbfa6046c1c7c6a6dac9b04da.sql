DELETE FROM system.webassembly_modules WHERE name = 'collatz';

-- Eliminar en masa todos los módulos cuyo nombre empiece por `tmp_` (el guion bajo literal se escapa como `\_`):
DELETE FROM system.webassembly_modules WHERE name LIKE 'tmp\_%';
