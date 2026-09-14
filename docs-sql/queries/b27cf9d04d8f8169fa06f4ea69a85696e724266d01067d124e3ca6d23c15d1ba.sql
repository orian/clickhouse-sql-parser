-- Encuentra todos los logs con datos de rate limiting:
SELECT * FROM logs WHERE has(mapKeys(attributes), 'rate_limit'); -- escaneo lento de tabla completa

-- Encuentra todos los logs de una IP específica:
SELECT * FROM logs WHERE has(mapValues(attributes), '192.168.1.1'); -- escaneo lento de tabla completa
