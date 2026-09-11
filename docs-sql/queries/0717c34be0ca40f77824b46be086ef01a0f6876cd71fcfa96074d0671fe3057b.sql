-- ✓ Valores en la cláusula WHERE
SELECT * FROM users WHERE id = {user_id: UInt64};

-- ✓ Nombres de tabla/base de datos
SELECT * FROM {db: Identifier}.{table: Identifier};

-- ✓ Valores en la cláusula IN
SELECT * FROM products WHERE id IN {ids: Array(UInt32)};

-- ✓ CREATE TABLE
CREATE TABLE {table_name: Identifier} (id UInt64, name String) ENGINE = MergeTree() ORDER BY id;
