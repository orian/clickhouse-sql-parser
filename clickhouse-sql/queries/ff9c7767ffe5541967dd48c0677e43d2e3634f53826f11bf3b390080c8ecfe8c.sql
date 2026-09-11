
        WITH 'xxxxyyyyxxxxyyyyxxxxyyyyxxxxyyyy' AS cipherKey 
        SELECT id, 
            AVG(value) OVER (ORDER BY id ROWS BETWEEN 1 PRECEDING AND 1 FOLLOWING) AS frame, 
            toString(frame) AS str,
            encrypt('aes-256-ofb', str, cipherKey) AS enc, 
            decrypt('aes-256-ofb', str, cipherKey) AS dec 
        FROM window_test
    