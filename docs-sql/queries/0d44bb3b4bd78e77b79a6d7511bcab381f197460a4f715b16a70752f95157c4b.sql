WHERE cityHash64(user_id) % 10 = 0  -- 항상 동일한 10%의 사용자만 선택
