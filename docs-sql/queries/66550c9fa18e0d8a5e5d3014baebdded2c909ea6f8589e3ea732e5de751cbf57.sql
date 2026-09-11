CREATE TABLE employees
(
    id UInt32,
    name String,
    salary UInt32
)
ENGINE = Memory;

INSERT INTO employees VALUES
    (87432, 'John Smith', 45680),
    (59018, 'Jane Smith', 72350),
    (20376, 'Ivan Ivanovich', 58900),
    (71245, 'Anastasia Ivanovna', 89210);

SELECT sum(salary) FROM employees;
