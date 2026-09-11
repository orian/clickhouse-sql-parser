-- 按类别设置不同的价格阈值
SELECT * FROM products
WHERE if(category = 'Electronics', price < 500, price < 200);
-- 结果：Mouse、Chair、Monitor
-- （Electronics 低于 $500 或 Furniture 低于 $200）

-- 根据库存状态过滤
SELECT * FROM products
WHERE if(in_stock, price > 100, true);
-- 结果：Laptop、Chair、Monitor、Desk、Lamp
-- （有库存且价格超过 $100 的商品，或所有缺货商品）
