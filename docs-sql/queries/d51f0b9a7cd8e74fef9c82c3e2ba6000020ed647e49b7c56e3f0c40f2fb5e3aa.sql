/* 原始查询 */
WITH 
  data AS (SELECT 1 AS id), 
  data AS (SELECT 2 AS id) -- 重新定义
SELECT * FROM data;

/* 修复后的查询 */
WITH 
  raw_data AS (SELECT 1 AS id), 
  processed_data AS (SELECT 2 AS id)
SELECT * FROM processed_data;
