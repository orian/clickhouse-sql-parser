-- カテゴリ別の価格しきい値
SELECT * FROM products
WHERE if(category = 'Electronics', price < 500, price < 200);
-- 結果: Mouse, Chair, Monitor
-- (Electronicsは$500未満、またはFurnitureは$200未満)

-- 在庫状況に基づくフィルター
SELECT * FROM products
WHERE if(in_stock, price > 100, true);
-- 結果: Laptop, Chair, Monitor, Desk, Lamp
-- ($100超の在庫あり商品、または在庫切れ商品すべて)
