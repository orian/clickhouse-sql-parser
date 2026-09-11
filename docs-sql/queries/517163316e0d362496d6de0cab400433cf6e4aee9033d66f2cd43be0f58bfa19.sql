CREATE TABLE salary
(
    user String,
    salary UInt32
)
ENGINE = Memory AS
SELECT *
FROM VALUES(('worker', 1000), ('manager', 3000), ('director', 5000));

SELECT argMax(user, salary) FROM salary;
