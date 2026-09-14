DROP TABLE IF EXISTS train_data;

CREATE TABLE train_data (target Float64, x1 Float64, x2 Float64) ENGINE = Memory;
INSERT INTO train_data VALUES (-1, 1, 1), (-1, 2, 1), (-1, 3, 2), (1, 8, 9), (1, 9, 8), (1, 10, 10);

DROP TABLE IF EXISTS your_model;

CREATE TABLE your_model
ENGINE = MergeTree
ORDER BY tuple()
AS SELECT
stochasticLogisticRegressionState(1.0, 1.0, 10, 'SGD')(target, x1, x2)
AS state FROM train_data;

DROP TABLE IF EXISTS test_data;

CREATE TABLE test_data (x1 Float64, x2 Float64) ENGINE = Memory;
INSERT INTO test_data VALUES (1, 1), (9, 9);

WITH (SELECT state FROM your_model) AS model
SELECT
evalMLMethod(model, x1, x2) BETWEEN 0 AND 1
FROM test_data
