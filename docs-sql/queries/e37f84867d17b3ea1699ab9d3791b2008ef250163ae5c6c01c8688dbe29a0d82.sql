WHERE cityHash64(user_id) % 10 = 0  -- Always same 10% of users
