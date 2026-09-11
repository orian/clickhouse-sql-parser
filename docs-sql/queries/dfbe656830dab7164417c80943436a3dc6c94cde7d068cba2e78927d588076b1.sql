CREATE TABLE user_shopping_sessions
(
    session_date Date,
    user_segment String,
    viewed_products Array(String),
    session_duration_minutes Int32
) ENGINE = Memory;

INSERT INTO user_shopping_sessions VALUES
    ('2024-01-01', 'new_customer', ['smartphone_x', 'headphones_y', 'smartphone_x'], 12),
    ('2024-01-01', 'returning', ['laptop_z', 'smartphone_x', 'tablet_a'], 25),
    ('2024-01-01', 'new_customer', ['smartwatch_b', 'headphones_y', 'fitness_tracker'], 8),
    ('2024-01-02', 'returning', ['laptop_z', 'external_drive', 'laptop_z'], 30),
    ('2024-01-02', 'new_customer', ['tablet_a', 'keyboard_c', 'tablet_a'], 15),
    ('2024-01-02', 'premium', ['smartphone_x', 'smartwatch_b', 'headphones_y'], 22);

-- 按细分类型和参与度级别统计不重复产品数量
SELECT 
    session_date,
    -- 统计新客户在长会话中浏览的不重复产品数量
    uniqArrayIf(viewed_products, user_segment = 'new_customer' AND session_duration_minutes > 10) AS new_customer_engaged_products,
    -- 统计回头客浏览的不重复产品数量
    uniqArrayIf(viewed_products, user_segment = 'returning') AS returning_customer_products,
    -- 统计所有会话中浏览的不重复产品数量
    uniqArray(viewed_products) AS total_unique_products
FROM user_shopping_sessions
GROUP BY session_date
ORDER BY session_date
FORMAT Vertical;
