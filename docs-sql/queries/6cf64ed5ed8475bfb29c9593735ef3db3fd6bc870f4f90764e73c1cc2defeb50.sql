CREATE TABLE employees
(
    id UInt32,
    name String,
    monthly_salary UInt16 -- selected so that the sum of values produces an overflow
)
ENGINE = Memory;

INSERT INTO employees VALUES
    (1, 'John', 20000),
    (2, 'Jane', 18000),
    (3, 'Bob', 12000),
    (4, 'Alice', 10000),
    (5, 'Charlie', 8000);

-- Query for the total amount of the employee salaries using the sum and sumWithOverflow functions and show their types using the toTypeName function
-- For the sum function the resulting type is UInt64, big enough to contain the sum, whilst for sumWithOverflow the resulting type remains as UInt16.

SELECT
    sum(monthly_salary) AS no_overflow,
    sumWithOverflow(monthly_salary) AS overflow,
    toTypeName(no_overflow),
    toTypeName(overflow)
FROM employees;
