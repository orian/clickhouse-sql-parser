/* 元のクエリ */
WITH 
  data AS (SELECT 1 AS id), 
  data AS (SELECT 2 AS id) -- 再定義
SELECT * FROM data;

/* 修正後のクエリ */
WITH 
  raw_data AS (SELECT 1 AS id), 
  processed_data AS (SELECT 2 AS id)
SELECT * FROM processed_data;
