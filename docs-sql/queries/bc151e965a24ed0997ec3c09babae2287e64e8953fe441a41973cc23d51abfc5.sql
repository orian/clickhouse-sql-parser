-- Ejecutar en origen y destino y comparar los resultados
SELECT MAX(id) AS latest_id, MAX(updated_at) AS latest_ts FROM public.orders;
