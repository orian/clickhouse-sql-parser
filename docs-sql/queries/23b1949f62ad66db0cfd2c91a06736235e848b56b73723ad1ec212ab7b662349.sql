-- Criar tabela de origem
CREATE TABLE source_data (
    id UInt32,
    name String,
    value Float64
) ENGINE = MergeTree
ORDER BY id;

-- Inserir alguns dados
INSERT INTO source_data VALUES (1, 'one', 10.1), (2, 'two', 20.2);

-- Criar alias
CREATE TABLE data_alias ENGINE = Alias('source_data');

-- Consultar via alias
SELECT * FROM data_alias;
