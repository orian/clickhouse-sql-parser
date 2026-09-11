WHERE cityHash64(user_id) % 10 = 0  -- Siempre el mismo 10% de usuarios
