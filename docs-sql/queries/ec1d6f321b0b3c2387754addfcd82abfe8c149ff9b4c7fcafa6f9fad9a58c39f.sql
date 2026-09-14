-- 카테고리별 다른 가격 임계값
SELECT * FROM products
WHERE if(category = 'Electronics', price < 500, price < 200);
-- 결과: Mouse, Chair, Monitor
-- (Electronics 중 $500 미만 OR Furniture 중 $200 미만)

-- 재고 상태에 따른 필터
SELECT * FROM products
WHERE if(in_stock, price > 100, true);
-- 결과: Laptop, Chair, Monitor, Desk, Lamp
-- ($100 초과 재고 보유 항목 OR 재고 없는 모든 항목)
