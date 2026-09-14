-- This MySQL-style prepared statement with parameters does NOT work in ClickHouse
PREPARE stmt FROM 'SELECT * FROM users WHERE id = ?';
EXECUTE stmt USING @user_id;  -- Parameter binding not supported
