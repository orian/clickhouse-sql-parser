-- 但当 key 或 iv 超过预期长度时，encrypt 会失败：
SELECT encrypt('aes-256-ofb', 'Secret', '123456789101213141516171819202122', 'iviviviviviviviv123');
