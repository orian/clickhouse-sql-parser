--SQL estándar
SELECT
   avg(number)
FILTER (WHERE number > 50)
FROM numbers(100)

--ClickHouse usando un combinador de agregación
SELECT
   avgIf(number, number > 50)
FROM numbers(100)
