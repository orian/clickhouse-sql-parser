-- Búsqueda sin distinción entre mayúsculas y minúsculas para 'LAPTOP'
SELECT * FROM products WHERE name ILIKE '%laptop%';
-- Resultado: Laptop

-- Coincidencia de prefijo sin distinción entre mayúsculas y minúsculas
SELECT * FROM products WHERE name ILIKE 'l%';
-- Resultado: Laptop, Lamp
