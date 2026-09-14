CREATE TABLE employees
(
    id UInt32,
    name String,
    monthly_salary UInt16 -- 选择此类型是为了让值的总和产生溢出
)
ENGINE = Memory;

INSERT INTO employees VALUES
    (1, 'John', 20000),
    (2, 'Jane', 18000),
    (3, 'Bob', 12000),
    (4, 'Alice', 10000),
    (5, 'Charlie', 8000);

-- 使用 sum 和 sumWithOverflow 函数查询员工薪资总额，并通过 toTypeName 函数显示其结果类型
-- sum 函数的结果类型为 UInt64，足以容纳总和；而 sumWithOverflow 的结果类型仍保持为 UInt16。

SELECT
    sum(monthly_salary) AS no_overflow,
    sumWithOverflow(monthly_salary) AS overflow,
    toTypeName(no_overflow),
    toTypeName(overflow)
FROM employees;
