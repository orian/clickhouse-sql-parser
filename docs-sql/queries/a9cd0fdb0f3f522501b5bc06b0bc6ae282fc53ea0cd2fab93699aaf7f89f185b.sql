-- 查找所有包含限流数据的日志：
SELECT * FROM logs WHERE has(mapKeys(attributes), 'rate_limit'); -- 全表扫描，速度较慢

-- 查找所有来自特定 IP 的日志：
SELECT * FROM logs WHERE has(mapValues(attributes), '192.168.1.1'); -- 全表扫描，速度较慢
