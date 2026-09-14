DELETE FROM system.webassembly_modules WHERE name = 'collatz';

-- Массовое удаление всех модулей, имя которых начинается с `tmp_` (символ подчёркивания экранируется как `\_`):
DELETE FROM system.webassembly_modules WHERE name LIKE 'tmp\_%';
