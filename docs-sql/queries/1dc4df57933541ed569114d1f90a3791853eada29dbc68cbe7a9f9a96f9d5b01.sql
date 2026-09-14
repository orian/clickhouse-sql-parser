CREATE TABLE employees (name String, salary UInt32) ENGINE = Memory;
INSERT INTO employees VALUES ('Alice', 3000), ('Bob', 4000), ('Charlie', 3500);

SELECT max(salary) FROM employees;
