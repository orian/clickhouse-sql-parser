-- 在输入相同的情况下，encrypt 和 aes_encrypt_mysql 会生成相同的密文：
SELECT encrypt('aes-256-ofb', 'Secret', '12345678910121314151617181920212', 'iviviviviviviviv') = aes_encrypt_mysql('aes-256-ofb', 'Secret', '12345678910121314151617181920212', 'iviviviviviviviv') AS ciphertexts_equal;
