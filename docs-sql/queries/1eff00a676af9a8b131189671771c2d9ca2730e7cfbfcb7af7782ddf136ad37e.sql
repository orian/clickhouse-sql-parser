CREATE TABLE users (
    id UInt32,
    name String
) ENGINE = MergeTree
ORDER BY id;

CREATE TABLE users_alias ENGINE = Alias('users');

-- 별칭을 통해 컬럼 추가
ALTER TABLE users_alias ADD COLUMN email String DEFAULT '';

-- 대상 테이블에 컬럼이 추가됨
DESCRIBE users;
