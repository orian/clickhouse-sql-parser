DELETE FROM system.webassembly_modules WHERE name = 'collatz';

-- 批量删除所有名称以 `tmp_` 开头的模块（字面下划线需转义为 `\_`）：
DELETE FROM system.webassembly_modules WHERE name LIKE 'tmp\_%';
