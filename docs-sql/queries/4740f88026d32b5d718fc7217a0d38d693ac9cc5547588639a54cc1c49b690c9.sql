-- Создать исходную таблицу
CREATE TABLE source_data (
    id UInt32,
    name String,
    value Float64
) ENGINE = MergeTree
ORDER BY id;

-- Вставить данные
INSERT INTO source_data VALUES (1, 'one', 10.1), (2, 'two', 20.2);

-- Создать псевдоним
CREATE TABLE data_alias ENGINE = Alias('source_data');

-- Выполнить запрос через псевдоним
SELECT * FROM data_alias;
