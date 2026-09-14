SELECT
    location_id,
    location_name,
    min_temp,     -- SimpleAggregateFunction の値に直接アクセス
    max_temp      -- SimpleAggregateFunction では集約完了関数は不要
FROM temperature_extremes
ORDER BY location_id;
