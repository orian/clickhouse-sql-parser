-- WRONG: compares the entire Array to a scalar
SELECT * FROM test.visits
WHERE Goals.ID = 591325;
