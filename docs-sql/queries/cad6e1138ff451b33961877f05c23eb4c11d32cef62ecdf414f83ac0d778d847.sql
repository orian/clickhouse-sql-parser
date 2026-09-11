-- 名前に 'o' を含む商品を検索
SELECT * FROM products WHERE name LIKE '%o%';
-- 結果: Laptop, Monitor

-- 'L' で始まる商品を検索
SELECT * FROM products WHERE name LIKE 'L%';
-- 結果: Laptop, Lamp

-- ちょうど4文字の商品を検索
SELECT * FROM products WHERE name LIKE '____';
-- 結果: Desk, Lamp
