CREATE TABLE users (
    id UInt32,
    name String
) ENGINE = MergeTree
ORDER BY id;

CREATE TABLE users_alias ENGINE = Alias('users');

-- 通过别名添加列
ALTER TABLE users_alias ADD COLUMN email String DEFAULT '';

-- 列已添加到目标表
DESCRIBE users;
