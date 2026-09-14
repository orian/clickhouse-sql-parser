-- 不区分大小写搜索 'LAPTOP'
SELECT * FROM products WHERE name ILIKE '%laptop%';
-- 结果：Laptop

-- 不区分大小写的前缀匹配
SELECT * FROM products WHERE name ILIKE 'l%';
-- 结果：Laptop, Lamp
