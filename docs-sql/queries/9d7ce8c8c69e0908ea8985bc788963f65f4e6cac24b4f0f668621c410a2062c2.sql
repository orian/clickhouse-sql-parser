CREATE TABLE IF NOT EXISTS test_data
(
    x1 Float64,
    x2 Float64
) ENGINE = Memory;

INSERT INTO test_data VALUES (10, 0), (20, 0);

WITH (SELECT state FROM your_model) AS model SELECT
evalMLMethod(model, x1, x2) FROM test_data
