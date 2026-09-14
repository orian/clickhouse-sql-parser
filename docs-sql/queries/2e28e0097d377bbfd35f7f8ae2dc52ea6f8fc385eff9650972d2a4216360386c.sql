-- Execute na origem e no destino e compare os resultados
SELECT MAX(id) AS latest_id, MAX(updated_at) AS latest_ts FROM public.orders;
