-- レート制限されたリクエストをすべて検索:
SELECT * FROM logs WHERE mapContainsKey(attributes, 'rate_limit'); -- fast

-- 特定のIPからのログをすべて検索:
SELECT * FROM logs WHERE has(mapValues(attributes), '192.168.1.1'); -- fast

-- いずれかの属性にエラーが含まれるログをすべて検索:
SELECT * FROM logs WHERE mapContainsValueLike(attributes, '% error %'); -- fast
