-- rate limitingデータを含むすべてのログを検索:
SELECT * FROM logs WHERE has(mapKeys(attributes), 'rate_limit'); -- 低速なフルテーブルスキャン

-- 特定のIPからのすべてのログを検索:
SELECT * FROM logs WHERE has(mapValues(attributes), '192.168.1.1'); -- 低速なフルテーブルスキャン
