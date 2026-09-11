-- Q2.4: Durante los últimos 3 meses, ¿qué clientes han realizado un número excesivo de solicitudes?

SELECT client_ip,
       COUNT(*) AS num_requests
FROM logs2
WHERE log_time >= TIMESTAMP '2012-10-01 00:00:00'
GROUP BY client_ip
HAVING COUNT(*) >= 100000
ORDER BY num_requests DESC;
