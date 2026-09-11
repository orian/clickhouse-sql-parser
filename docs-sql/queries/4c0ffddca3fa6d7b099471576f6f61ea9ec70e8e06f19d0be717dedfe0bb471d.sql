-- ✗ SELECT 中的列名（谨慎使用 Identifier）
SELECT {column: Identifier} FROM users;  -- 支持有限

-- ✗ 任意 SQL 片段
SELECT * FROM users {where_clause: String};  -- 不支持

-- ✗ ALTER TABLE 语句
ALTER TABLE {table: Identifier} ADD COLUMN new_col String;  -- 不支持

-- ✗ 多条语句
{statements: String};  -- 不支持
