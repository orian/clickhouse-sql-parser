--standard SQL
SELECT avg(DISTINCT number)

--ClickHouse using an aggregate combinator
SELECT avgDistinct(number)
