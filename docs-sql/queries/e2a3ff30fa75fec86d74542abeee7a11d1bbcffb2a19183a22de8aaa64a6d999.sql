-- 매개변수를 사용하여 뷰 쿼리
SELECT *
FROM sales_by_date(start_date='2024-01-01', end_date='2024-01-31')
WHERE product_id = 12345;
