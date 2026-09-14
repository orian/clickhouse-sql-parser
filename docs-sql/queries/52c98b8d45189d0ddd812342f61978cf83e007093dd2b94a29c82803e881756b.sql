SELECT * FROM products
WHERE and(or(category = 'Electronics', price > 100), in_stock);
