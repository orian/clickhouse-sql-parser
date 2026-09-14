-- Datos de muestra: el usuario visita el checkout -> detalle del producto -> checkout de nuevo -> pago
-- Sin allow_reentry: se detiene en el nivel 2 (página de detalle del producto)
-- Con allow_reentry: alcanza el nivel 4 (finalización del pago)

SELECT
    level,
    count() AS users
FROM
(
    SELECT
        user_id,
        windowFunnel(3600, 'strict_order', 'allow_reentry')(
            timestamp,
            action = 'begin_checkout',      -- Paso 1: Iniciar checkout
            action = 'view_product_detail', -- Paso 2: Ver detalle del producto
            action = 'begin_checkout',      -- Paso 3: Iniciar checkout de nuevo (reentrada)
            action = 'complete_payment'     -- Paso 4: Completar el pago
        ) AS level
    FROM user_events
    WHERE event_date = today()
    GROUP BY user_id
)
GROUP BY level
ORDER BY level ASC;
