-- ✗ SELECT のカラム名（Identifier は慎重に使用すること）
SELECT {column: Identifier} FROM users;  -- サポートが限定的

-- ✗ 任意の SQL フラグメント
SELECT * FROM users {where_clause: String};  -- 非サポート

-- ✗ ALTER TABLE ステートメント
ALTER TABLE {table: Identifier} ADD COLUMN new_col String;  -- 非サポート

-- ✗ 複数のステートメント
{statements: String};  -- 非サポート
