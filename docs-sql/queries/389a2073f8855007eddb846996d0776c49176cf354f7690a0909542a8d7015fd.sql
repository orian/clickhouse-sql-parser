-- ✓ WHERE 子句中的值
SELECT * FROM users WHERE id = {user_id: UInt64};

-- ✓ 表/数据库名称
SELECT * FROM {db: Identifier}.{table: Identifier};

-- ✓ IN 子句中的值
SELECT * FROM products WHERE id IN {ids: Array(UInt32)};

-- ✓ CREATE TABLE
CREATE TABLE {table_name: Identifier} (id UInt64, name String) ENGINE = MergeTree() ORDER BY id;
