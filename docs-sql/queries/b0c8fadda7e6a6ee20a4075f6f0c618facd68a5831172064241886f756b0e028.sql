-- Find visits that have at least one goal with ID greater than 1000000
SELECT CounterID, VisitID, Goals.ID
FROM test.visits
WHERE arrayExists(id -> id > 1000000, Goals.ID);
