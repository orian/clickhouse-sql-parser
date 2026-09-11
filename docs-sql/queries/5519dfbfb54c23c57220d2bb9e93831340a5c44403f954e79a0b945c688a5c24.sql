-- カテゴリに基づく複数の条件
SELECT * FROM products
WHERE multiIf(
    category = 'Electronics', price < 600,
    category = 'Furniture', in_stock = true,
    false
);
-- 結果: Mouse, Monitor, Chair
-- (Electronics < $600 または在庫ありの Furniture)

-- 段階的なフィルタリング
SELECT * FROM products
WHERE multiIf(
    price > 500, category = 'Electronics',
    price > 100, in_stock = true,
    true
);
-- 結果: Laptop, Chair, Monitor, Lamp
