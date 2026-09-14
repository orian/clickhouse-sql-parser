-- Reemplaza con tu tabla y columnas reales
SELECT 
    count() as total_rows,
    uniq(your_group_by_columns) as unique_combinations,
    round(uniq(your_group_by_columns) / count() * 100, 2) as aggregation_ratio
FROM your_table
WHERE your_filter_conditions;

-- Si aggregation_ratio > 70%, reconsidera el diseño de tu MV
-- Si aggregation_ratio < 10%, obtendrás una buena compresión
