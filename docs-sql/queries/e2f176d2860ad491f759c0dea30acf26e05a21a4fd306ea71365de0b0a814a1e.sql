/* ORIGINAL QUERY */
-- device_id is ambiguous
SELECT user_id, device_id FROM table GROUP BY user_id

/* FIXED QUERY */
SELECT user_id, any(device_id) FROM table GROUP BY user_id
-- OR
SELECT user_id, device_id FROM table GROUP BY user_id, device_id
