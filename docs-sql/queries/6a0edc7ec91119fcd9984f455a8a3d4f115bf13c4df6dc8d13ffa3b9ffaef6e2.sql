-- 重要なテーブルの行数比較
SELECT 'public.orders' AS table_name, COUNT(*) AS row_count FROM public.orders;
SELECT 'public.customers' AS table_name, COUNT(*) AS row_count FROM public.customers;

-- 高頻度テーブルの最新レコードの抜き取り確認
SELECT MAX(updated_at) FROM public.orders;
SELECT MAX(id) FROM public.orders;
