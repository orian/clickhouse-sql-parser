-- サンプルデータ: ユーザーがチェックアウト -> 商品詳細 -> 再度チェックアウト -> 支払いの順に操作
-- allow_reentryなし: レベル2で停止（商品詳細ページ）
-- allow_reentryあり: レベル4に到達（支払い完了）

SELECT
    level,
    count() AS users
FROM
(
    SELECT
        user_id,
        windowFunnel(3600, 'strict_order', 'allow_reentry')(
            timestamp,
            action = 'begin_checkout',      -- ステップ1: チェックアウト開始
            action = 'view_product_detail', -- ステップ2: 商品詳細を表示
            action = 'begin_checkout',      -- ステップ3: 再度チェックアウト開始（再エントリ）
            action = 'complete_payment'     -- ステップ4: 支払い完了
        ) AS level
    FROM user_events
    WHERE event_date = today()
    GROUP BY user_id
)
GROUP BY level
ORDER BY level ASC;
