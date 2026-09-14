-- Q3.1: A temperatura interna atingiu o ponto de congelamento durante o fim de semana?

SELECT *
FROM logs3
WHERE event_type = 'temperature'
  AND event_value <= 32.0
  AND log_time >= '2019-11-29 17:00:00.000';
