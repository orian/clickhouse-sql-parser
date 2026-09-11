--SQL estándar
SELECT avg(DISTINCT number)

--ClickHouse usando un combinador de agregación
SELECT avgDistinct(number)
