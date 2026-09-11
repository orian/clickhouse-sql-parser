-- Dados de exemplo: usuário visita checkout -> detalhe do produto -> checkout novamente -> pagamento
-- Sem allow_reentry: para no nível 2 (página de detalhe do produto)
-- Com allow_reentry: alcança o nível 4 (conclusão do pagamento)

SELECT
    level,
    count() AS users
FROM
(
    SELECT
        user_id,
        windowFunnel(3600, 'strict_order', 'allow_reentry')(
            timestamp,
            action = 'begin_checkout',      -- Passo 1: Iniciar checkout
            action = 'view_product_detail', -- Passo 2: Visualizar detalhe do produto  
            action = 'begin_checkout',      -- Passo 3: Iniciar checkout novamente (reentrada)
            action = 'complete_payment'     -- Passo 4: Concluir pagamento
        ) AS level
    FROM user_events
    WHERE event_date = today()
    GROUP BY user_id
)
GROUP BY level
ORDER BY level ASC;
