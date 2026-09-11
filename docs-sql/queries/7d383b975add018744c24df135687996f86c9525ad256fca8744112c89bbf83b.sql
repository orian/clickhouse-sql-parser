-- 실제 테이블과 컬럼으로 교체하십시오
SELECT 
    count() as total_rows,
    uniq(your_group_by_columns) as unique_combinations,
    round(uniq(your_group_by_columns) / count() * 100, 2) as aggregation_ratio
FROM your_table
WHERE your_filter_conditions;

-- aggregation_ratio가 70%를 초과하면 MV 설계를 재검토하십시오
-- aggregation_ratio가 10% 미만이면 압축 효과가 좋습니다
