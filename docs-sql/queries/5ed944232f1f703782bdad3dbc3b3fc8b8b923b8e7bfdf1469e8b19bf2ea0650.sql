--standard SQL
SELECT
   avg(number)
FILTER (WHERE number > 50)
FROM numbers(100)

--ClickHouse using an aggregate combinator
SELECT
   avgIf(number, number > 50)
FROM numbers(100)
