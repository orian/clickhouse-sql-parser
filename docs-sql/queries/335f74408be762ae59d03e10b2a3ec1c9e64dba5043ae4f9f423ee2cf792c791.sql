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

-- sum関数とsumWithOverflow関数を使用して従業員の給与合計を求め、toTypeName関数でその型を表示するクエリ
-- sum関数の結果型はUInt64であり合計値を格納するのに十分な大きさだが、sumWithOverflowの結果型はUInt16のままになる。

SELECT
    sum(monthly_salary) AS no_overflow,
    sumWithOverflow(monthly_salary) AS overflow,
    toTypeName(no_overflow),
    toTypeName(overflow)
FROM employees;
