-- 内存高效聚合模式：每个条件 = 每组一个整数
-- 核心思路：无论数据量多大，sumIf() 的内存占用始终有界
-- 每组内存占用：N 个整数（N * 8 字节），其中 N = 条件数量

SELECT 
    your_grouping_column,
    
    -- 每个 sumIf 在每组中仅创建一个整数计数器
    -- 无论每个条件匹配多少行，内存占用始终保持不变
    sumIf(1, your_condition_1) as condition_1_count,
    sumIf(1, your_condition_2) as condition_2_count,
    sumIf(1, your_text_column LIKE '%pattern%') as pattern_matches,
    sumIf(1, your_numeric_column > threshold_value) as above_threshold,
    
    -- 复杂的多条件聚合同样只占用固定大小的内存
    sumIf(1, your_condition_1 AND your_text_column LIKE '%pattern%') as complex_condition_count,
    
    -- 常规聚合（供参考）
    count() as total_rows,
    avg(your_numeric_column) as average_value,
    max(your_timestamp_column) as latest_timestamp
    
FROM your_schema.your_table
WHERE your_timestamp_column >= 'start_date' 
  AND your_timestamp_column < 'end_date'
GROUP BY your_grouping_column
HAVING condition_1_count > minimum_threshold 
   OR condition_2_count > another_threshold
ORDER BY (condition_1_count + condition_2_count + pattern_matches) DESC
LIMIT 20
