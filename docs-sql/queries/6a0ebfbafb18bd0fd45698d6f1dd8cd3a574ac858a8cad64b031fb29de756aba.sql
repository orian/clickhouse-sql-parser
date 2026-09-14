-- 이름에 'o'가 포함된 제품 찾기
SELECT * FROM products WHERE name LIKE '%o%';
-- 결과: Laptop, Monitor

-- 'L'로 시작하는 제품 찾기
SELECT * FROM products WHERE name LIKE 'L%';
-- 결과: Laptop, Lamp

-- 정확히 4글자인 제품 찾기
SELECT * FROM products WHERE name LIKE '____';
-- 결과: Desk, Lamp
