/* ИСХОДНЫЙ ЗАПРОС */
WITH 
  data AS (SELECT 1 AS id), 
  data AS (SELECT 2 AS id) -- Переопределено
SELECT * FROM data;

/* ИСПРАВЛЕННЫЙ ЗАПРОС */
WITH 
  raw_data AS (SELECT 1 AS id), 
  processed_data AS (SELECT 2 AS id)
SELECT * FROM processed_data;
