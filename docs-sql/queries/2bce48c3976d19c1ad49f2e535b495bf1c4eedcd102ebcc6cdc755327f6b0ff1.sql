-- ✓ WHERE 절의 값
SELECT * FROM users WHERE id = {user_id: UInt64};

-- ✓ 테이블/데이터베이스 이름
SELECT * FROM {db: Identifier}.{table: Identifier};

-- ✓ IN 절의 값
SELECT * FROM products WHERE id IN {ids: Array(UInt32)};

-- ✓ CREATE TABLE
CREATE TABLE {table_name: Identifier} (id UInt64, name String) ENGINE = MergeTree() ORDER BY id;
