CREATE TABLE employees
(
    id UInt32,
    name String,
    monthly_salary UInt16 -- выбрано так, чтобы сумма значений вызывала переполнение
)
ENGINE = Memory;

INSERT INTO employees VALUES
    (1, 'John', 20000),
    (2, 'Jane', 18000),
    (3, 'Bob', 12000),
    (4, 'Alice', 10000),
    (5, 'Charlie', 8000);

-- Запрос общей суммы зарплат сотрудников с использованием функций sum и sumWithOverflow, а также вывод их типов с помощью функции toTypeName
-- Для функции sum результирующий тип — UInt64, достаточно большой для хранения суммы, тогда как для sumWithOverflow результирующий тип остаётся UInt16.

SELECT
    sum(monthly_salary) AS no_overflow,
    sumWithOverflow(monthly_salary) AS overflow,
    toTypeName(no_overflow),
    toTypeName(overflow)
FROM employees;
