--SQL padrão
SELECT avg(DISTINCT number)

--ClickHouse usando um combinador de função agregada
SELECT avgDistinct(number)
