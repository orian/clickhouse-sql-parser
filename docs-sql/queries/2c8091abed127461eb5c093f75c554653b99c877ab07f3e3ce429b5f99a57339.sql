-- 关键表的行数对比
SELECT 'public.orders' AS table_name, COUNT(*) AS row_count FROM public.orders;
SELECT 'public.customers' AS table_name, COUNT(*) AS row_count FROM public.customers;

-- 抽查高活跃表中的最新记录
SELECT MAX(updated_at) FROM public.orders;
SELECT MAX(id) FROM public.orders;
