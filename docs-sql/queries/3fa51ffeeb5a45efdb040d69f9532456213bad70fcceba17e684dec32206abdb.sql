-- ✗ Имена столбцов в SELECT (используйте Identifier с осторожностью)
SELECT {column: Identifier} FROM users;  -- Ограниченная поддержка

-- ✗ Произвольные фрагменты SQL
SELECT * FROM users {where_clause: String};  -- НЕ ПОДДЕРЖИВАЕТСЯ

-- ✗ Команды ALTER TABLE
ALTER TABLE {table: Identifier} ADD COLUMN new_col String;  -- НЕ ПОДДЕРЖИВАЕТСЯ

-- ✗ Несколько команд
{statements: String};  -- НЕ ПОДДЕРЖИВАЕТСЯ
