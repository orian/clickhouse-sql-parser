-- Q2.1: Quais requisições causaram erros de servidor nas últimas 2 semanas?

SELECT *
FROM logs2
WHERE status_code >= 500
  AND log_time >= TIMESTAMP '2012-12-18 00:00:00'
ORDER BY log_time;
