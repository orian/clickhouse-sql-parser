-- 더 긴 IV를 제공해도 동일한 결과가 생성되는 것을 확인하세요
SELECT hex(aes_encrypt_mysql('aes-256-ofb', 'Secret', '123456789101213141516171819202122', 'iviviviviviviviv123456')) AS ciphertext
