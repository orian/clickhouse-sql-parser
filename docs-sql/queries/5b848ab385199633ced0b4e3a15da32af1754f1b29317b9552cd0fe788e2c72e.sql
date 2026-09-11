DELETE FROM system.webassembly_modules WHERE name = 'collatz';

-- Exclusão em massa de todos os módulos cujo nome começa com `tmp_` (o sublinhado literal é escapado como `\_`):
DELETE FROM system.webassembly_modules WHERE name LIKE 'tmp\_%';
