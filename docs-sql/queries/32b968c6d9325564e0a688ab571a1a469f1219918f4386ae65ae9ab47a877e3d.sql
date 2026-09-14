-- Регистронезависимый поиск по 'LAPTOP'
SELECT * FROM products WHERE name ILIKE '%laptop%';
-- Результат: Laptop

-- Регистронезависимое совпадение по префиксу
SELECT * FROM products WHERE name ILIKE 'l%';
-- Результат: Laptop, Lamp
