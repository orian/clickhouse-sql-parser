-- ✓ Values in WHERE clause
SELECT * FROM users WHERE id = {user_id: UInt64};

-- ✓ Table/database names
SELECT * FROM {db: Identifier}.{table: Identifier};

-- ✓ Values in IN clause
SELECT * FROM products WHERE id IN {ids: Array(UInt32)};

-- ✓ CREATE TABLE
CREATE TABLE {table_name: Identifier} (id UInt64, name String) ENGINE = MergeTree() ORDER BY id;
