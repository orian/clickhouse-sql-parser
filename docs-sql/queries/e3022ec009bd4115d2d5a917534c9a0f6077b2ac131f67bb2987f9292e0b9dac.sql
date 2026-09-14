CREATE TABLE employees
(
    id UInt32,
    name String,
    monthly_salary UInt16 -- selecionado para que a soma dos valores produza um overflow
)
ENGINE = Memory;

INSERT INTO employees VALUES
    (1, 'John', 20000),
    (2, 'Jane', 18000),
    (3, 'Bob', 12000),
    (4, 'Alice', 10000),
    (5, 'Charlie', 8000);

-- Consulta para o valor total dos salários dos funcionários usando as funções sum e sumWithOverflow e exibe seus tipos usando a função toTypeName
-- Para a função sum, o tipo resultante é UInt64, grande o suficiente para conter a soma, enquanto para sumWithOverflow o tipo resultante permanece UInt16.

SELECT
    sum(monthly_salary) AS no_overflow,
    sumWithOverflow(monthly_salary) AS overflow,
    toTypeName(no_overflow),
    toTypeName(overflow)
FROM employees;
