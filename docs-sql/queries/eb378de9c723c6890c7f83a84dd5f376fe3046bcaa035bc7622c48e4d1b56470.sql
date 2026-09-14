-- Замените на вашу реальную таблицу и столбцы
SELECT 
    count() as total_rows,
    uniq(your_group_by_columns) as unique_combinations,
    round(uniq(your_group_by_columns) / count() * 100, 2) as aggregation_ratio
FROM your_table
WHERE your_filter_conditions;

-- Если aggregation_ratio > 70%, пересмотрите дизайн вашего MV
-- Если aggregation_ratio < 10%, вы получите хорошее сжатие
