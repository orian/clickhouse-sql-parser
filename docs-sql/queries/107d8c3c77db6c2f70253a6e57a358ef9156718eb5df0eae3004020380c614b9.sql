CREATE TABLE sales (department String, revenue UInt32) ENGINE = Memory;
INSERT INTO sales VALUES ('Engineering', 100000), ('Engineering', 120000), ('Marketing', 80000), ('Marketing', 90000);

SELECT department, max(revenue) FROM sales GROUP BY department ORDER BY department;
