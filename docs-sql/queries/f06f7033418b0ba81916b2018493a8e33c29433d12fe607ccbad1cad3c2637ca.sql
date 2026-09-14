-- 하지만 key 또는 iv가 예상보다 길면 encrypt가 실패합니다:
SELECT encrypt('aes-256-ofb', 'Secret', '123456789101213141516171819202122', 'iviviviviviviviv123');
