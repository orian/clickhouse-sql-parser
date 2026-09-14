-- Buscar productos con 'o' en el nombre
SELECT * FROM products WHERE name LIKE '%o%';
-- Resultado: Laptop, Monitor

-- Buscar productos que empiecen por 'L'
SELECT * FROM products WHERE name LIKE 'L%';
-- Resultado: Laptop, Lamp

-- Buscar productos con exactamente 4 caracteres
SELECT * FROM products WHERE name LIKE '____';
-- Resultado: Desk, Lamp
