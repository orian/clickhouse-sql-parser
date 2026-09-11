SELECT type, count()
FROM system.documentation
GROUP BY type
ORDER BY count() DESC;
