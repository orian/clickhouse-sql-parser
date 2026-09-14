-- Пример данных: пользователь посещает оформление заказа -> страницу товара -> снова оформление заказа -> оплату
-- Без allow_reentry: останавливается на уровне 2 (страница товара)
-- С allow_reentry: достигает уровня 4 (завершение оплаты)

SELECT
    level,
    count() AS users
FROM
(
    SELECT
        user_id,
        windowFunnel(3600, 'strict_order', 'allow_reentry')(
            timestamp,
            action = 'begin_checkout',      -- Шаг 1: Начало оформления заказа
            action = 'view_product_detail', -- Шаг 2: Просмотр страницы товара
            action = 'begin_checkout',      -- Шаг 3: Повторное оформление заказа (повторный вход)
            action = 'complete_payment'     -- Шаг 4: Завершение оплаты
        ) AS level
    FROM user_events
    WHERE event_date = today()
    GROUP BY user_id
)
GROUP BY level
ORDER BY level ASC;
