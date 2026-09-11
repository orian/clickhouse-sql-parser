-- 챌린지: toStartOfMinute 또는 toStartOfWeek 같은 다양한 시간 함수를 사용해 보세요
-- 실험: 직접 보유한 timestamp 데이터로 카디널리티 차이를 비교해 보세요
SELECT 
    'Microsecond precision' as granularity,
    uniq(created_at) as unique_values,
    'Creates massive cardinality - bad for sort key' as impact
FROM github.github_events
WHERE created_at >= '2024-01-01'
UNION ALL
SELECT 
    'Hour precision',
    uniq(toStartOfHour(created_at)),
    'Much better for sort key - enables skip indexing'
FROM github.github_events
WHERE created_at >= '2024-01-01'
UNION ALL  
SELECT 
    'Day precision',
    uniq(toStartOfDay(created_at)),
    'Best for reporting queries'
FROM github.github_events
WHERE created_at >= '2024-01-01';
