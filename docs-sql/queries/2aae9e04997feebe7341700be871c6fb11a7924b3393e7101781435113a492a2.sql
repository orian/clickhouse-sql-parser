-- Esta sentencia preparada de estilo MySQL con parámetros NO funciona en ClickHouse
PREPARE stmt FROM 'SELECT * FROM users WHERE id = ?';
EXECUTE stmt USING @user_id;  -- Vinculación de parámetros no compatible
