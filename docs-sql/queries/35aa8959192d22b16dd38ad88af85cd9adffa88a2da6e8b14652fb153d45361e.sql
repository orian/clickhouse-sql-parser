/* 原始查询 */
-- device_id 存在歧义
SELECT user_id, device_id FROM table GROUP BY user_id

/* 修复后的查询 */
SELECT user_id, any(device_id) FROM table GROUP BY user_id
-- 或
SELECT user_id, device_id FROM table GROUP BY user_id, device_id
