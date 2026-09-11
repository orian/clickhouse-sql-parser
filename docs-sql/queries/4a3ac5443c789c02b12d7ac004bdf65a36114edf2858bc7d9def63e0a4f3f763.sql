-- Q2.1: ¿Qué solicitudes han causado errores del servidor en las últimas 2 semanas?

SELECT *
FROM logs2
WHERE status_code >= 500
  AND log_time >= TIMESTAMP '2012-12-18 00:00:00'
ORDER BY log_time;
