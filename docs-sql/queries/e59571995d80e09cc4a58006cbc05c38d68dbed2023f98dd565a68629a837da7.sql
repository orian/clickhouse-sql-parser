-- 查找所有受速率限制的请求：
SELECT * FROM logs WHERE mapContainsKey(attributes, 'rate_limit'); -- fast

-- 查找来自特定 IP 的所有日志：
SELECT * FROM logs WHERE has(mapValues(attributes), '192.168.1.1'); -- fast

-- 查找任意属性值包含错误信息的所有日志：
SELECT * FROM logs WHERE mapContainsValueLike(attributes, '% error %'); -- fast
