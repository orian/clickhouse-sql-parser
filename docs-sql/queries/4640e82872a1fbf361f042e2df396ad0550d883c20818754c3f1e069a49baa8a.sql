-- 메모리 효율적인 집계 패턴: 각 조건 = 그룹당 정수 하나
-- 핵심 원리: sumIf()는 데이터 볼륨에 관계없이 메모리 사용량을 일정 범위로 제한
-- 그룹당 메모리: N개의 정수 (N * 8 바이트), N = 조건 수

SELECT 
    your_grouping_column,
    
    -- 각 sumIf는 그룹당 정확히 하나의 정수 카운터를 생성
    -- 각 조건에 일치하는 행 수와 무관하게 메모리 사용량은 일정하게 유지
    sumIf(1, your_condition_1) as condition_1_count,
    sumIf(1, your_condition_2) as condition_2_count,
    sumIf(1, your_text_column LIKE '%pattern%') as pattern_matches,
    sumIf(1, your_numeric_column > threshold_value) as above_threshold,
    
    -- 복잡한 다중 조건 집계도 동일하게 일정한 메모리만 사용
    sumIf(1, your_condition_1 AND your_text_column LIKE '%pattern%') as complex_condition_count,
    
    -- 참고용 표준 집계
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
