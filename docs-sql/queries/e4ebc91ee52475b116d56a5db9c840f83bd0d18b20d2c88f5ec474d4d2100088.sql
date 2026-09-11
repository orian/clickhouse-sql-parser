/* CONSULTA ORIGINAL */
-- device_id es ambiguo
SELECT user_id, device_id FROM table GROUP BY user_id

/* CONSULTA CORREGIDA */
SELECT user_id, any(device_id) FROM table GROUP BY user_id
-- O
SELECT user_id, device_id FROM table GROUP BY user_id, device_id
