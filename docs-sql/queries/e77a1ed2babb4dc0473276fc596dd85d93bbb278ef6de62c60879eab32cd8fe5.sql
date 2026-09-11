-- Q2.4: 지난 3개월 동안 과도하게 많은 요청을 보낸 클라이언트는 무엇입니까?

SELECT client_ip,
       COUNT(*) AS num_requests
FROM logs2
WHERE log_time >= TIMESTAMP '2012-10-01 00:00:00'
GROUP BY client_ip
HAVING COUNT(*) >= 100000
ORDER BY num_requests DESC;
