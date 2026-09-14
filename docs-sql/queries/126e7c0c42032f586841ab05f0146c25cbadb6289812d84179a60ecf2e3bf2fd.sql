CREATE TABLE employees
(
    id UInt32,
    name String,
    monthly_salary UInt16 -- 값의 합산 결과가 오버플로우를 일으키도록 선택된 타입
)
ENGINE = Memory;

INSERT INTO employees VALUES
    (1, 'John', 20000),
    (2, 'Jane', 18000),
    (3, 'Bob', 12000),
    (4, 'Alice', 10000),
    (5, 'Charlie', 8000);

-- sum 함수와 sumWithOverflow 함수를 사용하여 직원 급여 총합을 쿼리하고, toTypeName 함수로 각 결과의 타입을 표시
-- sum 함수의 결과 타입은 합계를 충분히 수용할 수 있는 UInt64이며, sumWithOverflow의 결과 타입은 UInt16으로 유지됨

SELECT
    sum(monthly_salary) AS no_overflow,
    sumWithOverflow(monthly_salary) AS overflow,
    toTypeName(no_overflow),
    toTypeName(overflow)
FROM employees;
