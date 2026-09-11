-- Q2.5: 일별 고유 방문자 수는 얼마입니까?

SELECT dt,
       COUNT(DISTINCT client_ip)
FROM (
  SELECT CAST(log_time AS DATE) AS dt,
         client_ip
  FROM logs2
) AS r
GROUP BY dt
ORDER BY dt;
