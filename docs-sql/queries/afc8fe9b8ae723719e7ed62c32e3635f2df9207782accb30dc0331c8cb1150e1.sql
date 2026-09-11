CREATE TABLE foo.randomNumbers
ORDER BY number AS
SELECT rand() AS number
FROM numbers(10_000_000);
