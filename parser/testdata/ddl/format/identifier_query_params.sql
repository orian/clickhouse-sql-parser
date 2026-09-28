-- Origin SQL:
CREATE DATABASE {db:Identifier};
CREATE TABLE {db:Identifier}.t (x Int8) ENGINE = Memory;
ALTER TABLE {db:Identifier}.t ADD COLUMN y Int8;
TRUNCATE TABLE {t:Identifier};
DROP TABLE {db:Identifier}.t;
DROP DATABASE IF EXISTS {db:Identifier};


-- Format SQL:
CREATE DATABASE {db: Identifier};
CREATE TABLE {db: Identifier}.t (x Int8) ENGINE = Memory;
ALTER TABLE {db: Identifier}.t ADD COLUMN y Int8;
TRUNCATE TABLE {t: Identifier};
DROP TABLE {db: Identifier}.t;
DROP DATABASE IF EXISTS {db: Identifier};
