-- Crear bases de datos
CREATE DATABASE db1;
CREATE DATABASE db2;

-- Crear tabla fuente en db1
CREATE TABLE db1.events (
    timestamp DateTime,
    event_type String,
    user_id UInt32
) ENGINE = MergeTree
ORDER BY timestamp;

-- Crear alias en db2 que apunta a db1.events
CREATE TABLE db2.events_alias ENGINE = Alias('db1', 'events');

-- O usando el formato base_de_datos.tabla
CREATE TABLE db2.events_alias2 ENGINE = Alias('db1.events');

-- Ambos alias funcionan de forma idéntica
INSERT INTO db2.events_alias VALUES (now(), 'click', 100);
SELECT * FROM db2.events_alias2;
