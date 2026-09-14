DROP TABLE IF EXISTS train_data;

CREATE TABLE train_data (target Float64, x1 Float64, x2 Float64) ENGINE = Memory;
INSERT INTO train_data VALUES (1, 1, 0), (2, 2, 0), (3, 3, 0), (4, 4, 0), (5, 5, 0), (6, 6, 0);

DROP TABLE IF EXISTS your_model;

CREATE TABLE your_model
ENGINE = Memory
AS SELECT
stochasticLinearRegressionState(0.1, 0.0, 5, 'SGD')(target, x1, x2)
AS state FROM train_data;

DROP TABLE IF EXISTS test_data;

CREATE TABLE test_data (x1 Float64, x2 Float64) ENGINE = Memory;
INSERT INTO test_data VALUES (10, 0), (20, 0);

WITH (SELECT state FROM your_model) AS model SELECT
evalMLMethod(model, x1, x2) > 0 FROM test_data
