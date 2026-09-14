SELECT name, content
FROM filesystem('my_directory')
WHERE name LIKE '%.csv';
