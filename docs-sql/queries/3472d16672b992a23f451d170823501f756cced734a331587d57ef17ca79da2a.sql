/* 元のクエリ */
-- device_id は曖昧です
SELECT user_id, device_id FROM table GROUP BY user_id

/* 修正後のクエリ */
SELECT user_id, any(device_id) FROM table GROUP BY user_id
-- または
SELECT user_id, device_id FROM table GROUP BY user_id, device_id
