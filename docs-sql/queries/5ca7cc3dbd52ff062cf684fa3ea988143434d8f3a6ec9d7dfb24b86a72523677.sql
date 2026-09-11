-- Without the filter, the mapped IPv4 addresses show up as well.
SELECT
    IPv6NumToString(ClientIP6 AS k),
    count() AS c
FROM hits_all
WHERE EventDate = today()
GROUP BY k
ORDER BY c DESC, k
LIMIT 10
