-- メモリ効率の高い集計パターン: 各条件 = グループごとに1つの整数
-- ポイント: sumIf() はデータ量に関わらずメモリ使用量を一定に抑える
-- グループごとのメモリ: N個の整数 (N * 8 バイト)、N = 条件の数

SELECT 
    your_grouping_column,
    
    -- 各 sumIf はグループごとに整数カウンターを1つだけ生成する
    -- 各条件に一致する行数に関わらず、メモリ使用量は一定
    sumIf(1, your_condition_1) as condition_1_count,
    sumIf(1, your_condition_2) as condition_2_count,
    sumIf(1, your_text_column LIKE '%pattern%') as pattern_matches,
    sumIf(1, your_numeric_column > threshold_value) as above_threshold,
    
    -- 複雑な複数条件の集計でもメモリ使用量は一定
    sumIf(1, your_condition_1 AND your_text_column LIKE '%pattern%') as complex_condition_count,
    
    -- 参考用の標準集計
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
