SELECT name, value, changed 
FROM system.errors 
WHERE value > 0 
ORDER BY value DESC;
