DROP TABLE IF EXISTS test_data;
CREATE TABLE test_data
(
    x UInt8,
)
ENGINE = Memory;

INSERT INTO test_data VALUES (3), (3), (3), (4), (4), (5), (5), (7), (11), (15);

SELECT
    varPop(x) AS var_pop
FROM test_data;
