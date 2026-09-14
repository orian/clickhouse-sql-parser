WHERE cityHash64(user_id) % 10 = 0  -- 始终抽取相同的 10% 用户
