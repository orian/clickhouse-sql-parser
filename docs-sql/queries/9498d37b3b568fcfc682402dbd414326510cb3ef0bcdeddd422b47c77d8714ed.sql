-- 여러 카테고리 기반 조건
SELECT * FROM products
WHERE multiIf(
    category = 'Electronics', price < 600,
    category = 'Furniture', in_stock = true,
    false
);
-- 결과: Mouse, Monitor, Chair
-- (Electronics < $600 또는 재고 있는 Furniture)

-- 단계별 필터링
SELECT * FROM products
WHERE multiIf(
    price > 500, category = 'Electronics',
    price > 100, in_stock = true,
    true
);
-- 결과: Laptop, Chair, Monitor, Lamp
