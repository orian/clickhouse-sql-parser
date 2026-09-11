-- 查找名称中包含 'o' 的产品
SELECT * FROM products WHERE name LIKE '%o%';
-- 结果：Laptop, Monitor

-- 查找名称以 'L' 开头的产品
SELECT * FROM products WHERE name LIKE 'L%';
-- 结果：Laptop, Lamp

-- 查找名称恰好为 4 个字符的产品
SELECT * FROM products WHERE name LIKE '____';
-- 结果：Desk, Lamp
