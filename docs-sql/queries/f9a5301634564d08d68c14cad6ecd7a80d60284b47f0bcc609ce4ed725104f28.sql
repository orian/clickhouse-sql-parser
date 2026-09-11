-- Обратите внимание: передача ещё более длинного IV даёт тот же результат
SELECT hex(aes_encrypt_mysql('aes-256-ofb', 'Secret', '123456789101213141516171819202122', 'iviviviviviviviv123456')) AS ciphertext
