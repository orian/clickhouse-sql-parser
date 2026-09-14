-- 実際のテーブルとカラムに置き換えてください
SELECT 
    count() as total_rows,
    uniq(your_group_by_columns) as unique_combinations,
    round(uniq(your_group_by_columns) / count() * 100, 2) as aggregation_ratio
FROM your_table
WHERE your_filter_conditions;

-- aggregation_ratio が 70% を超える場合は、materialized view の設計を見直してください
-- aggregation_ratio が 10% 未満の場合は、高い圧縮率が期待できます
