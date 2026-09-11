-- Let's decrypt data we've previously encrypted with MySQL:
-- mysql> SET  block_encryption_mode='aes-256-ofb';
-- Query OK, 0 rows affected (0.00 sec)
--
-- mysql> SELECT aes_encrypt('Secret', '123456789101213141516171819202122', 'iviviviviviviviv123456') as ciphertext;
-- +------------------------+
-- | ciphertext             |
-- +------------------------+
-- | 0x24E9E4966469         |
-- +------------------------+
-- 1 row in set (0.00 sec)

SELECT aes_decrypt_mysql('aes-256-ofb', unhex('24E9E4966469'), '123456789101213141516171819202122', 'iviviviviviviviv123456') AS plaintext
