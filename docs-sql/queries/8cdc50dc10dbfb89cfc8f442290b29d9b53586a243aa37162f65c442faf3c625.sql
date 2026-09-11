-- Padrão de Agregação com Uso Eficiente de Memória: Cada condição = um inteiro por grupo
-- Ponto-chave: sumIf() cria memória limitada independentemente do volume de dados
-- Memória por grupo: N inteiros (N * 8 bytes) onde N = número de condições

SELECT 
    your_grouping_column,
    
    -- Cada sumIf cria exatamente um contador inteiro por grupo
    -- A memória permanece constante independentemente de quantas linhas satisfazem cada condição
    sumIf(1, your_condition_1) as condition_1_count,
    sumIf(1, your_condition_2) as condition_2_count,
    sumIf(1, your_text_column LIKE '%pattern%') as pattern_matches,
    sumIf(1, your_numeric_column > threshold_value) as above_threshold,
    
    -- Agregações complexas com múltiplas condições ainda utilizam memória constante
    sumIf(1, your_condition_1 AND your_text_column LIKE '%pattern%') as complex_condition_count,
    
    -- Agregações padrão para referência
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
