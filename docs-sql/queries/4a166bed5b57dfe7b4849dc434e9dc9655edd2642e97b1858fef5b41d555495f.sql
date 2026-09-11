SELECT
    location_id,
    location_name,
    min_temp,     -- SimpleAggregateFunction 값에 직접 접근
    max_temp      -- SimpleAggregateFunction은 finalization 함수가 필요 없음
FROM temperature_extremes
ORDER BY location_id;
