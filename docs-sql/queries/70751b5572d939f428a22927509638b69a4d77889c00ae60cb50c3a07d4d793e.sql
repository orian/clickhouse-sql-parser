-- 大文字と小文字を区別しない 'LAPTOP' の検索
SELECT * FROM products WHERE name ILIKE '%laptop%';
-- 結果: Laptop

-- 大文字と小文字を区別しないプレフィックス一致
SELECT * FROM products WHERE name ILIKE 'l%';
-- 結果: Laptop, Lamp
