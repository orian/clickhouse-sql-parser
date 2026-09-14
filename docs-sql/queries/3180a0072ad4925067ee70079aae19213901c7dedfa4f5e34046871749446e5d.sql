-- Этот подготовленный оператор в стиле MySQL с параметрами НЕ работает в ClickHouse
PREPARE stmt FROM 'SELECT * FROM users WHERE id = ?';
EXECUTE stmt USING @user_id;  -- Привязка параметров не поддерживается
