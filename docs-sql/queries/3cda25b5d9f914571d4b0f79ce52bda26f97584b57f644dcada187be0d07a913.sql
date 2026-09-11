-- ✗ SELECT의 컬럼 이름 (Identifier 사용 시 주의 필요)
SELECT {column: Identifier} FROM users;  -- 제한적 지원

-- ✗ 임의의 SQL 단편
SELECT * FROM users {where_clause: String};  -- 지원되지 않음

-- ✗ ALTER TABLE 문
ALTER TABLE {table: Identifier} ADD COLUMN new_col String;  -- 지원되지 않음

-- ✗ 다중 SQL 문
{statements: String};  -- 지원되지 않음
