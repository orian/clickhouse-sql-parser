-- 在源端和目标端分别运行并比较结果
SELECT MAX(id) AS latest_id, MAX(updated_at) AS latest_ts FROM public.orders;
