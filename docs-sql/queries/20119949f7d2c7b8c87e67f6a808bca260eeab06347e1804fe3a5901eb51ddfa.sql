-- But encrypt fails when key or iv is longer than expected:
SELECT encrypt('aes-256-ofb', 'Secret', '123456789101213141516171819202122', 'iviviviviviviviv123');
