-- Este prepared statement no estilo MySQL com parâmetros NÃO funciona no ClickHouse
PREPARE stmt FROM 'SELECT * FROM users WHERE id = ?';
EXECUTE stmt USING @user_id;  -- Vinculação de parâmetros não suportada
