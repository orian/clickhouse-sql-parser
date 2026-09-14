-- Busca sem distinção entre maiúsculas e minúsculas por 'LAPTOP'
SELECT * FROM products WHERE name ILIKE '%laptop%';
-- Result: Laptop

-- Correspondência de prefixo sem distinção entre maiúsculas e minúsculas
SELECT * FROM products WHERE name ILIKE 'l%';
-- Result: Laptop, Lamp
