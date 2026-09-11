-- Crear tabla fuente
CREATE TABLE source_data (
    id UInt32,
    name String,
    value Float64
) ENGINE = MergeTree
ORDER BY id;

-- Insertar algunos datos
INSERT INTO source_data VALUES (1, 'one', 10.1), (2, 'two', 20.2);

-- Crear alias
CREATE TABLE data_alias ENGINE = Alias('source_data');

-- Consultar a través del alias
SELECT * FROM data_alias;
