SELECT PostId, concat('https://stackoverflow.com/questions/', PostId) AS url, count() AS c
FROM votes
GROUP BY PostId
ORDER BY c DESC
LIMIT 5
