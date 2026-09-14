SELECT path, name, size
FROM filesystem()
WHERE type = 'regular' AND size > 1000000
ORDER BY size DESC;
