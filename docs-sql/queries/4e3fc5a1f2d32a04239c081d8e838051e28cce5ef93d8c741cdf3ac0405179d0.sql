-- ✓ Значения в предложении WHERE
SELECT * FROM users WHERE id = {user_id: UInt64};

-- ✓ Имена таблиц/баз данных
SELECT * FROM {db: Identifier}.{table: Identifier};

-- ✓ Значения в предложении IN
SELECT * FROM products WHERE id IN {ids: Array(UInt32)};

-- ✓ CREATE TABLE
CREATE TABLE {table_name: Identifier} (id UInt64, name String) ENGINE = MergeTree() ORDER BY id;
