-- Encontrar produtos com 'o' no nome
SELECT * FROM products WHERE name LIKE '%o%';
-- Resultado: Laptop, Monitor

-- Encontrar produtos que começam com 'L'
SELECT * FROM products WHERE name LIKE 'L%';
-- Resultado: Laptop, Lamp

-- Encontrar produtos com exatamente 4 caracteres
SELECT * FROM products WHERE name LIKE '____';
-- Resultado: Desk, Lamp
