/* ИСХОДНЫЙ ЗАПРОС */
-- device_id неоднозначен
SELECT user_id, device_id FROM table GROUP BY user_id

/* ИСПРАВЛЕННЫЙ ЗАПРОС */
SELECT user_id, any(device_id) FROM table GROUP BY user_id
-- ИЛИ
SELECT user_id, device_id FROM table GROUP BY user_id, device_id
