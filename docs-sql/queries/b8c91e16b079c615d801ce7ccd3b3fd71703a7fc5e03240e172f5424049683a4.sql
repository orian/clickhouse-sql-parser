-- Criar bancos de dados
CREATE DATABASE db1;
CREATE DATABASE db2;

-- Criar tabela de origem em db1
CREATE TABLE db1.events (
    timestamp DateTime,
    event_type String,
    user_id UInt32
) ENGINE = MergeTree
ORDER BY timestamp;

-- Criar alias em db2 apontando para db1.events
CREATE TABLE db2.events_alias ENGINE = Alias('db1', 'events');

-- Ou usando o formato database.table
CREATE TABLE db2.events_alias2 ENGINE = Alias('db1.events');

-- Ambos os aliases funcionam de forma idêntica
INSERT INTO db2.events_alias VALUES (now(), 'click', 100);
SELECT * FROM db2.events_alias2;
