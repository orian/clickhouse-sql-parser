-- Удаляет все строки из таблицы `hits`, где столбец `Title` содержит текст `hello`
DELETE FROM hits WHERE Title LIKE '%hello%';
