-- Comparação de contagem de linhas para tabelas críticas
SELECT 'public.orders' AS table_name, COUNT(*) AS row_count FROM public.orders;
SELECT 'public.customers' AS table_name, COUNT(*) AS row_count FROM public.customers;

-- Verificação rápida dos registros mais recentes em tabelas de alta atividade
SELECT MAX(updated_at) FROM public.orders;
SELECT MAX(id) FROM public.orders;
