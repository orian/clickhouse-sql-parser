-- ソースとターゲットの両方で実行し、結果を比較する
SELECT MAX(id) AS latest_id, MAX(updated_at) AS latest_ts FROM public.orders;
