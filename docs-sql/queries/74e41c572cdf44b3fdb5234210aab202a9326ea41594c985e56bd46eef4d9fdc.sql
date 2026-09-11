-- このMySQL形式のパラメータ付きプリペアドステートメントはClickHouseでは動作しません
PREPARE stmt FROM 'SELECT * FROM users WHERE id = ?';
EXECUTE stmt USING @user_id;  -- パラメータバインディングはサポートされていません
