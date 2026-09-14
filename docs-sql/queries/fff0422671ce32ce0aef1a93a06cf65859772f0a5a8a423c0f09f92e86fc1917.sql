CREATE TABLE ontime (AirlineID UInt32) ENGINE = Memory;

-- Three quarters of the flights belong to one airline, so it occurs in more than half of the
-- rows seen by every thread and the result is the same on every run.
INSERT INTO ontime SELECT if(number % 4 = 0, 19393, 19690) FROM numbers(1000);

SELECT anyHeavy(AirlineID) AS res
FROM ontime;
