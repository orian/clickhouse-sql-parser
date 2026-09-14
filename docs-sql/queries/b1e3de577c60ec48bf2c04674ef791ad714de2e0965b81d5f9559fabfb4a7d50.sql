-- Q2.1: 지난 2주 동안 서버 오류를 유발한 요청은 무엇입니까?

SELECT *
FROM logs2
WHERE status_code >= 500
  AND log_time >= TIMESTAMP '2012-12-18 00:00:00'
ORDER BY log_time;
