/* 원본 쿼리 */
WITH 
  data AS (SELECT 1 AS id), 
  data AS (SELECT 2 AS id) -- 재정의됨
SELECT * FROM data;

/* 수정된 쿼리 */
WITH 
  raw_data AS (SELECT 1 AS id), 
  processed_data AS (SELECT 2 AS id)
SELECT * FROM processed_data;
