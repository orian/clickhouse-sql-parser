CREATE TABLE users (
    id UInt32,
    name String
) ENGINE = MergeTree
ORDER BY id;

CREATE TABLE users_alias ENGINE = Alias('users');

-- Добавление столбца через псевдоним
ALTER TABLE users_alias ADD COLUMN email String DEFAULT '';

-- Столбец добавляется в целевую таблицу
DESCRIBE users;
