-- Сравнение количества строк в критически важных таблицах
SELECT 'public.orders' AS table_name, COUNT(*) AS row_count FROM public.orders;
SELECT 'public.customers' AS table_name, COUNT(*) AS row_count FROM public.customers;

-- Выборочная проверка последних записей в таблицах с высокой активностью
SELECT MAX(updated_at) FROM public.orders;
SELECT MAX(id) FROM public.orders;
