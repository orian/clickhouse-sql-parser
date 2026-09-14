-- Encontra todos os logs com dados de rate limiting:
SELECT * FROM logs WHERE has(mapKeys(attributes), 'rate_limit'); -- varredura lenta na tabela completa

-- Encontra todos os logs de um IP específico:
SELECT * FROM logs WHERE has(mapValues(attributes), '192.168.1.1'); -- varredura lenta na tabela completa
