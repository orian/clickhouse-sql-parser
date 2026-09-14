-- Несколько условий на основе категорий
SELECT * FROM products
WHERE multiIf(
    category = 'Electronics', price < 600,
    category = 'Furniture', in_stock = true,
    false
);
-- Результат: Mouse, Monitor, Chair
-- (Electronics < $600 ИЛИ Furniture в наличии)

-- Многоуровневая фильтрация
SELECT * FROM products
WHERE multiIf(
    price > 500, category = 'Electronics',
    price > 100, in_stock = true,
    true
);
-- Результат: Laptop, Chair, Monitor, Lamp
