-- Múltiplas condições baseadas em categoria
SELECT * FROM products
WHERE multiIf(
    category = 'Electronics', price < 600,
    category = 'Furniture', in_stock = true,
    false
);
-- Resultado: Mouse, Monitor, Chair
-- (Electronics < $600 OU Furniture em estoque)

-- Filtragem em camadas
SELECT * FROM products
WHERE multiIf(
    price > 500, category = 'Electronics',
    price > 100, in_stock = true,
    true
);
-- Resultado: Laptop, Chair, Monitor, Lamp
