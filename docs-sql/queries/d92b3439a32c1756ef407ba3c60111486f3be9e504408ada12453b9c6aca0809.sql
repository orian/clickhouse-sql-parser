-- Q2.1: 过去 2 周内哪些请求导致了服务器错误？

SELECT *
FROM logs2
WHERE status_code >= 500
  AND log_time >= TIMESTAMP '2012-12-18 00:00:00'
ORDER BY log_time;
