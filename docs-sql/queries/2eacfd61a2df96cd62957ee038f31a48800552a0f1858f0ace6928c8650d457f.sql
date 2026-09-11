-- Comparación de recuento de filas para tablas críticas
SELECT 'public.orders' AS table_name, COUNT(*) AS row_count FROM public.orders;
SELECT 'public.customers' AS table_name, COUNT(*) AS row_count FROM public.customers;

-- Verificación rápida de los registros más recientes en tablas de alta actividad
SELECT MAX(updated_at) FROM public.orders;
SELECT MAX(id) FROM public.orders;
