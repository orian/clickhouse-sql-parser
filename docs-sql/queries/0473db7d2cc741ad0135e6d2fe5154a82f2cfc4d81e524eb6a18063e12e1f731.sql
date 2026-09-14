-- Encontra todas as requisições com limite de taxa:
SELECT * FROM logs WHERE mapContainsKey(attributes, 'rate_limit'); -- fast

-- Encontra todos os logs de um IP específico:
SELECT * FROM logs WHERE has(mapValues(attributes), '192.168.1.1'); -- fast

-- Encontra todos os logs em que algum atributo inclui um erro:
SELECT * FROM logs WHERE mapContainsValueLike(attributes, '% error %'); -- fast
