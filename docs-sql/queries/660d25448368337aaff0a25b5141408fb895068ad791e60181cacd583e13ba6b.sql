/* 원본 쿼리 */
-- device_id가 모호함
SELECT user_id, device_id FROM table GROUP BY user_id

/* 수정된 쿼리 */
SELECT user_id, any(device_id) FROM table GROUP BY user_id
-- 또는
SELECT user_id, device_id FROM table GROUP BY user_id, device_id
