-- Find visits that have at least one goal with ID 591325
SELECT CounterID, VisitID, Goals.ID
FROM test.visits
WHERE has(Goals.ID, 591325);
