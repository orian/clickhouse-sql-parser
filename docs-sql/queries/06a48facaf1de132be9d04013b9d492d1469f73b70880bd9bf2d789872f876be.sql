-- Mas encrypt falha quando key ou iv é maior que o esperado:
SELECT encrypt('aes-256-ofb', 'Secret', '123456789101213141516171819202122', 'iviviviviviviviv123');
