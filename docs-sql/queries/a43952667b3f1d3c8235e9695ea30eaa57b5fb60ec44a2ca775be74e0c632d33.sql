-- Run on both source and target and compare results
SELECT MAX(id) AS latest_id, MAX(updated_at) AS latest_ts FROM public.orders;
