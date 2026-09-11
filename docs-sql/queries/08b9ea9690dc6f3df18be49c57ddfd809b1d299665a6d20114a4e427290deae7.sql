-- Различные пороговые значения цен по категориям
SELECT * FROM products
WHERE if(category = 'Electronics', price < 500, price < 200);
-- Результат: Mouse, Chair, Monitor
-- (Electronics дешевле $500 ИЛИ Furniture дешевле $200)

-- Фильтрация по статусу наличия на складе
SELECT * FROM products
WHERE if(in_stock, price > 100, true);
-- Результат: Laptop, Chair, Monitor, Desk, Lamp
-- (Товары в наличии дороже $100 ИЛИ все товары не в наличии)
