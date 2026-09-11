--SQL padrão
SELECT
   avg(number)
FILTER (WHERE number > 50)
FROM numbers(100)

--ClickHouse usando um combinador de função agregada
SELECT
   avgIf(number, number > 50)
FROM numbers(100)
