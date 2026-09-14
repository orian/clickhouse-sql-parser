-- 동일한 입력값이 주어지면 encrypt와 aes_encrypt_mysql은 동일한 암호문(ciphertext)을 생성합니다:
SELECT encrypt('aes-256-ofb', 'Secret', '12345678910121314151617181920212', 'iviviviviviviviv') = aes_encrypt_mysql('aes-256-ofb', 'Secret', '12345678910121314151617181920212', 'iviviviviviviviv') AS ciphertexts_equal;
