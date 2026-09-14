SELECT *
FROM products
WHERE xor(price > 200, category = 'Electronics')
