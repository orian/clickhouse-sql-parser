WHERE cityHash64(user_id) % 10 = 0  -- Всегда одни и те же 10% пользователей
