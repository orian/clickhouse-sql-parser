-- Encuentra todas las solicitudes con límite de tasa:
SELECT * FROM logs WHERE mapContainsKey(attributes, 'rate_limit'); -- rápido

-- Encuentra todos los logs de una IP específica:
SELECT * FROM logs WHERE has(mapValues(attributes), '192.168.1.1'); -- rápido

-- Encuentra todos los logs donde algún atributo incluye un error:
SELECT * FROM logs WHERE mapContainsValueLike(attributes, '% error %'); -- rápido
