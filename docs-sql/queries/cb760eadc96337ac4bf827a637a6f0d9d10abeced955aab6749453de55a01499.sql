-- Находит все журналы с данными об ограничении частоты запросов:
SELECT * FROM logs WHERE has(mapKeys(attributes), 'rate_limit'); -- медленное полное сканирование таблицы

-- Находит все журналы с определённого IP-адреса:
SELECT * FROM logs WHERE has(mapValues(attributes), '192.168.1.1'); -- медленное полное сканирование таблицы
