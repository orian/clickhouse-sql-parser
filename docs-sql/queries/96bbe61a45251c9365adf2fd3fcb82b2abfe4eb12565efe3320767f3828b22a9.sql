-- 删除 `hits` 表中 `Title` 列包含文本 `hello` 的所有行
DELETE FROM hits WHERE Title LIKE '%hello%';
