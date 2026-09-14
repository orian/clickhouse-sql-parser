-- Row count comparison for critical tables
SELECT 'public.orders' AS table_name, COUNT(*) AS row_count FROM public.orders;
SELECT 'public.customers' AS table_name, COUNT(*) AS row_count FROM public.customers;

-- Spot-check latest records in high-activity tables
SELECT MAX(updated_at) FROM public.orders;
SELECT MAX(id) FROM public.orders;
