  -- ClickHouse 네이티브 쿼리 매개변수(권장)
  SET param_user_id = 12345;
  SELECT * FROM users WHERE id = {user_id: UInt64};
