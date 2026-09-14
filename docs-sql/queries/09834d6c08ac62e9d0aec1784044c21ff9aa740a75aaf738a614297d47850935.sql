--표준 SQL
SELECT
   avg(number)
FILTER (WHERE number > 50)
FROM numbers(100)

--집계 조합자를 사용하는 ClickHouse
SELECT
   avgIf(number, number > 50)
FROM numbers(100)
