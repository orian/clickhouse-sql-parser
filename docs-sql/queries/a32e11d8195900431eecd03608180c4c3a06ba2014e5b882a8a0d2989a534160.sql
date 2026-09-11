-- Elimina todas las filas de la tabla `hits` donde la columna `Title` contiene el texto `hello`
DELETE FROM hits WHERE Title LIKE '%hello%';
