  -- معلمات الاستعلام الأصلية في ClickHouse (موصى بها)
  SET param_user_id = 12345;
  SELECT * FROM users WHERE id = {user_id: UInt64};
