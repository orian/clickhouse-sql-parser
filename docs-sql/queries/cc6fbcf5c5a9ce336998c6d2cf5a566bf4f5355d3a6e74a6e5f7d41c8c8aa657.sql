CREATE TABLE my_test_table
ENGINE = MergeTree
ORDER BY tuple()
AS SELECT * 
FROM generateRandom(
    'col1 UInt32, col2 String, col3 Float64, col4 DateTime',
    1,  -- seed для генерации данных
    10  -- количество различных случайных значений
)
LIMIT 100;  -- 100 строк

-- Шаг 2: запрос к новой таблице
SELECT * FROM my_test_table LIMIT 5;
