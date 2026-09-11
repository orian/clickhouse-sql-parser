-- Найти продукты с 'o' в названии
SELECT * FROM products WHERE name LIKE '%o%';
-- Результат: Laptop, Monitor

-- Найти продукты, начинающиеся с 'L'
SELECT * FROM products WHERE name LIKE 'L%';
-- Результат: Laptop, Lamp

-- Найти продукты, содержащие ровно 4 символа в названии
SELECT * FROM products WHERE name LIKE '____';
-- Результат: Desk, Lamp
