SELECT PostId, count() AS c
FROM postlinks
GROUP BY PostId
ORDER BY c DESC LIMIT 5
