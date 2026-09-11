SELECT count() FROM posts WHERE has(keywords, 'clickhouse'); -- медленное полное сканирование таблицы — проверяет каждое ключевое слово в каждом посте
