/* ORIGINAL QUERY */
WITH 
  data AS (SELECT 1 AS id), 
  data AS (SELECT 2 AS id) -- Redefined
SELECT * FROM data;

/* FIXED QUERY */
WITH 
  raw_data AS (SELECT 1 AS id), 
  processed_data AS (SELECT 2 AS id)
SELECT * FROM processed_data;
