-- Q3.1: ¿Alcanzó la temperatura interior el punto de congelación durante el fin de semana?

SELECT *
FROM logs3
WHERE event_type = 'temperature'
  AND event_value <= 32.0
  AND log_time >= '2019-11-29 17:00:00.000';
