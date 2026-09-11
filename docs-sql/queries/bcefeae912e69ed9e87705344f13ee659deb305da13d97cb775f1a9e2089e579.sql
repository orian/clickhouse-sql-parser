-- しかし、keyまたはivが期待される長さより長い場合、encryptは失敗する：
SELECT encrypt('aes-256-ofb', 'Secret', '123456789101213141516171819202122', 'iviviviviviviviv123');
