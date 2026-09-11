-- 간단한 계산용 함수
CREATE FUNCTION calculate_tax AS (price, rate) -> price * rate;

SELECT
    product_name,
    price,
    calculate_tax(price, 0.08) AS tax
FROM products;
