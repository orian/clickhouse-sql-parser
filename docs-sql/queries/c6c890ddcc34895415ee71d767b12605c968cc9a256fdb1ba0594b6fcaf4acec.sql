-- Diferentes umbrales de precio por categoría
SELECT * FROM products
WHERE if(category = 'Electronics', price < 500, price < 200);
-- Resultado: Mouse, Chair, Monitor
-- (Electronics por debajo de $500 O Furniture por debajo de $200)

-- Filtro según el estado del inventario
SELECT * FROM products
WHERE if(in_stock, price > 100, true);
-- Resultado: Laptop, Chair, Monitor, Desk, Lamp
-- (Artículos en stock por encima de $100 O todos los artículos sin stock)
