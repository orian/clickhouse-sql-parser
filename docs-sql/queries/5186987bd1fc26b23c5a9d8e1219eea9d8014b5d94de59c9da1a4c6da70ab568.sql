-- Создать базы данных
CREATE DATABASE db1;
CREATE DATABASE db2;

-- Создать исходную таблицу в db1
CREATE TABLE db1.events (
    timestamp DateTime,
    event_type String,
    user_id UInt32
) ENGINE = MergeTree
ORDER BY timestamp;

-- Создать псевдоним в db2, указывающий на db1.events
CREATE TABLE db2.events_alias ENGINE = Alias('db1', 'events');

-- Или в формате database.table
CREATE TABLE db2.events_alias2 ENGINE = Alias('db1.events');

-- Оба псевдонима работают одинаково
INSERT INTO db2.events_alias VALUES (now(), 'click', 100);
SELECT * FROM db2.events_alias2;
