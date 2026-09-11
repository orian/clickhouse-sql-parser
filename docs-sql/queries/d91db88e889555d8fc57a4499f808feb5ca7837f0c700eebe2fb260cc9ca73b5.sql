CREATE TABLE users (
    id UInt32,
    name String
) ENGINE = MergeTree
ORDER BY id;

CREATE TABLE users_alias ENGINE = Alias('users');

-- Adicionar coluna através do alias
ALTER TABLE users_alias ADD COLUMN email String DEFAULT '';

-- Coluna é adicionada à tabela de destino
DESCRIBE users;
