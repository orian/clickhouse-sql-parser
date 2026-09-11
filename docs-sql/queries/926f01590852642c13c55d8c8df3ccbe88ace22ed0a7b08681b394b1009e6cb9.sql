-- Replace with your actual table and columns
SELECT 
    count() as total_rows,
    uniq(your_group_by_columns) as unique_combinations,
    round(uniq(your_group_by_columns) / count() * 100, 2) as aggregation_ratio
FROM your_table
WHERE your_filter_conditions;

-- If aggregation_ratio > 70%, reconsider your MV design
-- If aggregation_ratio < 10%, you'll get good compression
