SELECT path, sum(hits) AS h
FROM wikistat
WHERE (project = 'it') AND (subproject = 'zero')
GROUP BY path
ORDER BY h DESC
LIMIT 10;
