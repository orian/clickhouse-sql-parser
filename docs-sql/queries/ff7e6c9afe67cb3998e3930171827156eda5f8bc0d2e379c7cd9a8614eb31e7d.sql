-- 替换为您的实际表和列
SELECT 
    count() as total_rows,
    uniq(your_group_by_columns) as unique_combinations,
    round(uniq(your_group_by_columns) / count() * 100, 2) as aggregation_ratio
FROM your_table
WHERE your_filter_conditions;

-- 如果 aggregation_ratio > 70%，请重新考虑物化视图的设计
-- 如果 aggregation_ratio < 10%，可获得良好的压缩效果
