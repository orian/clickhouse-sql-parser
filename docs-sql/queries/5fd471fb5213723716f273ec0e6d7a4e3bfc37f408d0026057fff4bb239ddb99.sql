-- Visits with at least 3 goals
SELECT CounterID, VisitID, Goals.ID
FROM test.visits
WHERE length(Goals.ID) >= 3;

-- Visits with at least one goal (non-empty array)
SELECT CounterID, VisitID, Goals.ID
FROM test.visits
WHERE notEmpty(Goals.ID);
