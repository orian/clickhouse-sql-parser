SELECT name, type
FROM filesystem('my_directory')
WHERE depth = 0;
