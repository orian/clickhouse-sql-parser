-- 基于多个类别的条件
SELECT * FROM products
WHERE multiIf(
    category = 'Electronics', price < 600,
    category = 'Furniture', in_stock = true,
    false
);
-- 结果：Mouse, Monitor, Chair
-- （Electronics 价格 < $600 或有库存的 Furniture）

-- 分层过滤
SELECT * FROM products
WHERE multiIf(
    price > 500, category = 'Electronics',
    price > 100, in_stock = true,
    true
);
-- 结果：Laptop, Chair, Monitor, Lamp
