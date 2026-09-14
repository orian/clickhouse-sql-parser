CREATE TABLE users (
    id UInt32,
    name String
) ENGINE = MergeTree
ORDER BY id;

CREATE TABLE users_alias ENGINE = Alias('users');

-- Agregar columna a través del alias
ALTER TABLE users_alias ADD COLUMN email String DEFAULT '';

-- La columna se agrega a la tabla de destino
DESCRIBE users;
