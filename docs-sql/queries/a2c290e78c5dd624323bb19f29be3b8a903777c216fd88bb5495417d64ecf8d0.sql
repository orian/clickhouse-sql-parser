WHERE cityHash64(user_id) % 10 = 0  -- Sempre os mesmos 10% dos usuários
