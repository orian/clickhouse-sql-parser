-- ✓ WHERE句の値
SELECT * FROM users WHERE id = {user_id: UInt64};

-- ✓ テーブル/データベース名
SELECT * FROM {db: Identifier}.{table: Identifier};

-- ✓ IN句の値
SELECT * FROM products WHERE id IN {ids: Array(UInt32)};

-- ✓ CREATE TABLE
CREATE TABLE {table_name: Identifier} (id UInt64, name String) ENGINE = MergeTree() ORDER BY id;
