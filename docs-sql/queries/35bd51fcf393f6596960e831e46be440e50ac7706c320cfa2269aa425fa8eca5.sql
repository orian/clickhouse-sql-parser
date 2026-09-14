-- Substitua pela sua tabela e colunas reais
SELECT 
    count() as total_rows,
    uniq(your_group_by_columns) as unique_combinations,
    round(uniq(your_group_by_columns) / count() * 100, 2) as aggregation_ratio
FROM your_table
WHERE your_filter_conditions;

-- Se aggregation_ratio > 70%, reconsidere o design da sua MV
-- Se aggregation_ratio < 10%, você terá boa compressão
