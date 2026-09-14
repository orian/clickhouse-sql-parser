--标准 SQL
SELECT
   avg(number)
FILTER (WHERE number > 50)
FROM numbers(100)

--ClickHouse 使用聚合组合器
SELECT
   avgIf(number, number > 50)
FROM numbers(100)
