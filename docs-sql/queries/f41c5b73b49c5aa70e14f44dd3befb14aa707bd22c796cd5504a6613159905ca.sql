-- 'LAPTOP'에 대한 대소문자를 구분하지 않는 검색
SELECT * FROM products WHERE name ILIKE '%laptop%';
-- 결과: Laptop

-- 대소문자를 구분하지 않는 접두사 일치
SELECT * FROM products WHERE name ILIKE 'l%';
-- 결과: Laptop, Lamp
