-- 使用 if() 的条件逻辑
CREATE FUNCTION price_tier AS (price) ->
    if(price < 100, 'Budget',
       if(price < 500, 'Mid-range', 'Premium'));

SELECT
    product_name,
    price,
    price_tier(price) AS tier
FROM products;
