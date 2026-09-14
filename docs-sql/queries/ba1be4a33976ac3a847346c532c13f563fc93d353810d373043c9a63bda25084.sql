DROP TABLE IF EXISTS train_data;

CREATE TABLE train_data (target Float64, x1 Float64, x2 Float64) ENGINE = Memory;
INSERT INTO train_data VALUES (1, 1, 0), (2, 2, 0), (3, 3, 0), (4, 4, 0), (5, 5, 0), (6, 6, 0);

SELECT length(stochasticLinearRegression(0.01)(target, x1, x2)) FROM train_data
