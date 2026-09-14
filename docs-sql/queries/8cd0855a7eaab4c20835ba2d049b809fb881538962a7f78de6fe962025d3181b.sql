-- Шаблон агрегации с экономным использованием памяти: каждое условие = одно целое число на группу
-- Ключевая идея: sumIf() ограничивает потребление памяти независимо от объёма данных
-- Память на группу: N целых чисел (N * 8 байт), где N = количество условий

SELECT 
    your_grouping_column,
    
    -- Каждый sumIf создаёт ровно один целочисленный счётчик на группу
    -- Объём памяти остаётся постоянным вне зависимости от того, сколько строк соответствует каждому условию
    sumIf(1, your_condition_1) as condition_1_count,
    sumIf(1, your_condition_2) as condition_2_count,
    sumIf(1, your_text_column LIKE '%pattern%') as pattern_matches,
    sumIf(1, your_numeric_column > threshold_value) as above_threshold,
    
    -- Сложные агрегации с несколькими условиями также используют постоянный объём памяти
    sumIf(1, your_condition_1 AND your_text_column LIKE '%pattern%') as complex_condition_count,
    
    -- Стандартные агрегации для справки
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
