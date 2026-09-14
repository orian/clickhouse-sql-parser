-- Diferentes limites de preço por categoria
SELECT * FROM products
WHERE if(category = 'Electronics', price < 500, price < 200);
-- Resultado: Mouse, Chair, Monitor
-- (Electronics abaixo de $500 OU Furniture abaixo de $200)

-- Filtrar com base no status do estoque
SELECT * FROM products
WHERE if(in_stock, price > 100, true);
-- Resultado: Laptop, Chair, Monitor, Desk, Lamp
-- (Itens em estoque acima de $100 OU todos os itens fora de estoque)
