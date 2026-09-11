--标准 SQL
SELECT avg(DISTINCT number)

--ClickHouse 使用聚合组合器
SELECT avgDistinct(number)
