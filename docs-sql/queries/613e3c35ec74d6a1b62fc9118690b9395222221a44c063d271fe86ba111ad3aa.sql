-- Each goal ID has a calculation of the number of conversions (each element in the Goals nested data structure is a goal that was reached, which we refer to as a conversion)
-- and the number of sessions. Without ARRAY JOIN, we would have counted the number of sessions as sum(Sign). But in this particular case,
-- the rows were multiplied by the nested Goals structure, so in order to count each session one time after this, we apply a condition to the
-- value of the arrayEnumerateUniq(Goals.ID) function.

CREATE TABLE visits (CounterID UInt32, Sign Int8, Goals Nested(ID UInt32)) ENGINE = Memory;

INSERT INTO visits VALUES
    (160656, 1, [53225, 53225, 56600]),
    (160656, 1, [53225, 2825062]),
    (160656, 1, [56600, 56600, 2825062]),
    (160657, 1, [53225]);

SELECT
    Goals.ID AS GoalID,
    sum(Sign) AS Reaches,
    sumIf(Sign, num = 1) AS Visits
FROM visits
ARRAY JOIN
    Goals,
    arrayEnumerateUniq(Goals.ID) AS num
WHERE CounterID = 160656
GROUP BY GoalID
ORDER BY Reaches DESC, GoalID
LIMIT 10
