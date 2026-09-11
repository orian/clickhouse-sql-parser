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

-- Contar produtos únicos visualizados por tipo de segmento e nível de engajamento
SELECT 
    session_date,
    -- Contar produtos únicos visualizados em sessões longas por novos clientes
    uniqArrayIf(viewed_products, user_segment = 'new_customer' AND session_duration_minutes > 10) AS new_customer_engaged_products,
    -- Contar produtos únicos visualizados por clientes recorrentes
    uniqArrayIf(viewed_products, user_segment = 'returning') AS returning_customer_products,
    -- Contar produtos únicos visualizados em todas as sessões
    uniqArray(viewed_products) AS total_unique_products
FROM user_shopping_sessions
GROUP BY session_date
ORDER BY session_date
FORMAT Vertical;
