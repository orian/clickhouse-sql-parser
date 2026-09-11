-- Найти все запросы с ограничением частоты:
SELECT * FROM logs WHERE mapContainsKey(attributes, 'rate_limit'); -- fast

-- Найти все записи журнала с определённого IP-адреса:
SELECT * FROM logs WHERE has(mapValues(attributes), '192.168.1.1'); -- fast

-- Найти все записи журнала, в которых любой атрибут содержит ошибку:
SELECT * FROM logs WHERE mapContainsValueLike(attributes, '% error %'); -- fast
