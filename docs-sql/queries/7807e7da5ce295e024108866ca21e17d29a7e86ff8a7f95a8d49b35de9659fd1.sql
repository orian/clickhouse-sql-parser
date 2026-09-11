-- Но encrypt завершается ошибкой, если key или iv длиннее ожидаемого:
SELECT encrypt('aes-256-ofb', 'Secret', '123456789101213141516171819202122', 'iviviviviviviviv123');
