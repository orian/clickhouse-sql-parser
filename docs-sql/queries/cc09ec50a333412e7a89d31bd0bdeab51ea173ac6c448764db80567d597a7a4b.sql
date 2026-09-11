SELECT status, request.method[1] AS method, count() AS c
FROM http
WHERE status >= 400
  AND toDateTime(timestamp) BETWEEN '1998-01-01 00:00:00' AND '1998-06-01 00:00:00'
GROUP BY method, status
ORDER BY c DESC LIMIT 5;
