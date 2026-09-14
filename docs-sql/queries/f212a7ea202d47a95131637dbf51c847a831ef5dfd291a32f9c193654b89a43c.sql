CREATE TABLE users (
    id UInt32,
    name String
) ENGINE = MergeTree
ORDER BY id;

CREATE TABLE users_alias ENGINE = Alias('users');

-- エイリアス経由でカラムを追加
ALTER TABLE users_alias ADD COLUMN email String DEFAULT '';

-- カラムがターゲットテーブルに追加される
DESCRIBE users;
