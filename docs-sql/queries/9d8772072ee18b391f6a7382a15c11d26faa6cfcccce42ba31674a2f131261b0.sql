-- 중요 테이블의 행 수 비교
SELECT 'public.orders' AS table_name, COUNT(*) AS row_count FROM public.orders;
SELECT 'public.customers' AS table_name, COUNT(*) AS row_count FROM public.customers;

-- 활동이 많은 테이블의 최신 레코드 점검
SELECT MAX(updated_at) FROM public.orders;
SELECT MAX(id) FROM public.orders;
