-- 소스와 타겟 모두에서 실행하여 결과를 비교하세요
SELECT MAX(id) AS latest_id, MAX(updated_at) AS latest_ts FROM public.orders;
