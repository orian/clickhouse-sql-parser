-- Q2.1: Какие запросы вызвали ошибки сервера за последние 2 недели?

SELECT *
FROM logs2
WHERE status_code >= 500
  AND log_time >= TIMESTAMP '2012-12-18 00:00:00'
ORDER BY log_time;
