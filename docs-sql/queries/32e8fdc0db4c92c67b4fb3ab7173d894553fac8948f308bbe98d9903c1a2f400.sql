CREATE TABLE my_test_table
ENGINE = MergeTree
ORDER BY tuple()
AS SELECT * 
FROM generateRandom(
    'col1 UInt32, col2 String, col3 Float64, col4 DateTime',
    1,  -- semilla para la generación de datos
    10  -- número de valores aleatorios distintos
)
LIMIT 100;  -- 100 filas

-- Paso 2: Consultar la nueva tabla
SELECT * FROM my_test_table LIMIT 5;
