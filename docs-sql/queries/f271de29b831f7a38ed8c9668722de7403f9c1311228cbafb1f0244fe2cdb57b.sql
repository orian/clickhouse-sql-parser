-- 这种带参数的 MySQL 风格预处理语句在 ClickHouse 中不起作用
PREPARE stmt FROM 'SELECT * FROM users WHERE id = ?';
EXECUTE stmt USING @user_id;  -- 不支持参数绑定
