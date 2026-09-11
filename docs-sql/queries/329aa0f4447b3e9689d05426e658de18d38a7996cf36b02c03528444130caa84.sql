--標準SQL
SELECT
   avg(number)
FILTER (WHERE number > 50)
FROM numbers(100)

--集約コンビネータを使用したClickHouse
SELECT
   avgIf(number, number > 50)
FROM numbers(100)
