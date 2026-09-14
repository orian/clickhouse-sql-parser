-- Condiciones múltiples basadas en categoría
SELECT * FROM products
WHERE multiIf(
    category = 'Electronics', price < 600,
    category = 'Furniture', in_stock = true,
    false
);
-- Resultado: Mouse, Monitor, Chair
-- (Electronics < $600 OR Furniture en stock)

-- Filtrado por niveles
SELECT * FROM products
WHERE multiIf(
    price > 500, category = 'Electronics',
    price > 100, in_stock = true,
    true
);
-- Resultado: Laptop, Chair, Monitor, Lamp
