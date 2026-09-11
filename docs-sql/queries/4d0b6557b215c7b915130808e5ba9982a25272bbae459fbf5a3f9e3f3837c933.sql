WHERE cityHash64(user_id) % 10 = 0  -- 常に同じ10%のユーザー
