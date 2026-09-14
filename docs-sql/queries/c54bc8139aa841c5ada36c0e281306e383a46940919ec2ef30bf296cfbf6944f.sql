CREATE TABLE employees
(
    id UInt32,
    name String,
    monthly_salary UInt16 -- seleccionado para que la suma de valores produzca un desbordamiento
)
ENGINE = Memory;

INSERT INTO employees VALUES
    (1, 'John', 20000),
    (2, 'Jane', 18000),
    (3, 'Bob', 12000),
    (4, 'Alice', 10000),
    (5, 'Charlie', 8000);

-- Consulta del importe total de los salarios de los empleados usando las funciones sum y sumWithOverflow, y muestra sus tipos con la función toTypeName
-- Para la función sum, el tipo resultante es UInt64, suficientemente grande para contener la suma, mientras que para sumWithOverflow el tipo resultante permanece como UInt16.

SELECT
    sum(monthly_salary) AS no_overflow,
    sumWithOverflow(monthly_salary) AS overflow,
    toTypeName(no_overflow),
    toTypeName(overflow)
FROM employees;
