/* CONSULTA ORIGINAL */
-- device_id é ambíguo
SELECT user_id, device_id FROM table GROUP BY user_id

/* CONSULTA CORRIGIDA */
SELECT user_id, any(device_id) FROM table GROUP BY user_id
-- OU
SELECT user_id, device_id FROM table GROUP BY user_id, device_id
