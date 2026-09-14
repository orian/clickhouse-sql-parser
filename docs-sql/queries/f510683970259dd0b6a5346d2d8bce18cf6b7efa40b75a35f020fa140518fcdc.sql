-- 样本数据：用户访问结账 -> 商品详情 -> 再次结账 -> 支付
-- 不使用 allow_reentry：停在第 2 级（商品详情页）
-- 使用 allow_reentry：到达第 4 级（完成支付）

SELECT
    level,
    count() AS users
FROM
(
    SELECT
        user_id,
        windowFunnel(3600, 'strict_order', 'allow_reentry')(
            timestamp,
            action = 'begin_checkout',      -- 第 1 步：开始结账
            action = 'view_product_detail', -- 第 2 步：查看商品详情
            action = 'begin_checkout',      -- 第 3 步：再次开始结账（重入）
            action = 'complete_payment'     -- 第 4 步：完成支付
        ) AS level
    FROM user_events
    WHERE event_date = today()
    GROUP BY user_id
)
GROUP BY level
ORDER BY level ASC;
