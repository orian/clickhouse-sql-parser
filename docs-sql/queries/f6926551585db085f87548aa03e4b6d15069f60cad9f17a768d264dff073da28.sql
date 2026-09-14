-- MySQL 스타일의 매개변수가 포함된 prepared statement는 ClickHouse에서 작동하지 않습니다
PREPARE stmt FROM 'SELECT * FROM users WHERE id = ?';
EXECUTE stmt USING @user_id;  -- 매개변수 바인딩은 지원되지 않습니다
