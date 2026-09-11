-- Выполните на источнике и целевой базе данных и сравните результаты
SELECT MAX(id) AS latest_id, MAX(updated_at) AS latest_ts FROM public.orders;
