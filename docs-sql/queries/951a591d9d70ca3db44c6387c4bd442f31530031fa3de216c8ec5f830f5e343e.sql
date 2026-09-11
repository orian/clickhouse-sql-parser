-- Q1.4: En 1 mes, ¿con qué frecuencia estuvo bloqueado cada servidor por I/O de disco?

SELECT machine_name,
       COUNT(*) AS spikes
FROM logs1
WHERE machine_group = 'Servers'
  AND cpu_wio > 0.99
  AND log_time >= TIMESTAMP '2016-12-01 00:00:00'
  AND log_time < TIMESTAMP '2017-01-01 00:00:00'
GROUP BY machine_name
ORDER BY spikes DESC
LIMIT 10;
