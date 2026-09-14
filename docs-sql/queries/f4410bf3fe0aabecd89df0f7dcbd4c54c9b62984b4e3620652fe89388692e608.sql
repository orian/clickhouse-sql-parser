-- Patrón de agregación eficiente en memoria: cada condición = un entero por grupo
-- Idea clave: sumIf() acota el uso de memoria independientemente del volumen de datos
-- Memoria por grupo: N enteros (N * 8 bytes) donde N = número de condiciones

SELECT 
    your_grouping_column,
    
    -- Cada sumIf crea exactamente un contador entero por grupo
    -- La memoria se mantiene constante sin importar cuántas filas coincidan con cada condición
    sumIf(1, your_condition_1) as condition_1_count,
    sumIf(1, your_condition_2) as condition_2_count,
    sumIf(1, your_text_column LIKE '%pattern%') as pattern_matches,
    sumIf(1, your_numeric_column > threshold_value) as above_threshold,
    
    -- Las agregaciones complejas con múltiples condiciones también usan memoria constante
    sumIf(1, your_condition_1 AND your_text_column LIKE '%pattern%') as complex_condition_count,
    
    -- Agregaciones estándar de referencia
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
