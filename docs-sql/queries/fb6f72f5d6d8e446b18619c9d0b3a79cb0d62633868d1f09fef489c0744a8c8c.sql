CREATE TABLE my_test_table
ENGINE = MergeTree
ORDER BY tuple()
AS SELECT * 
FROM generateRandom(
    'col1 UInt32, col2 String, col3 Float64, col4 DateTime',
    1,  -- semente para gerar os dados
    10  -- número de valores aleatórios diferentes
)
LIMIT 100;  -- 100 linhas

-- Etapa 2: consulte a nova tabela
SELECT * FROM my_test_table LIMIT 5;
