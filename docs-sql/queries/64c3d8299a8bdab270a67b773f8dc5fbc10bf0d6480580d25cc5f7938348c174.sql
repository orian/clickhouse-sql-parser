--стандартный SQL
SELECT
   avg(number)
FILTER (WHERE number > 50)
FROM numbers(100)

--ClickHouse с использованием комбинатора агрегатных функций
SELECT
   avgIf(number, number > 50)
FROM numbers(100)
