SELECT result < 1.1 AND result > 0.5 FROM
(WITH (SELECT state FROM your_model) AS model SELECT
evalMLMethod(model, x1, x2) AS result FROM test_data)
