--стандартный SQL
SELECT avg(DISTINCT number)

--ClickHouse с использованием комбинатора агрегатных функций
SELECT avgDistinct(number)
